package metrics

import (
	"net"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/achushu/libs/out"
	"github.com/achushu/libs/conv"
	gom "github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"
)

// Wrappers for go-metrics
type (
	// Registry is a metrics registry
	Registry gom.Registry
)

var (
	startTime         time.Time
	metricStartTime   gom.Gauge
	metricCurrentTime gom.Gauge
	metricUptime      gom.Gauge

	metricGoRoutines  gom.Gauge
	metricGC          gom.Gauge
	metricHeapAlloc   gom.Gauge
	metricHeapInUse   gom.Gauge
	metricHeapObjects gom.Gauge
	metricStackInUse gom.Gauge
	metricSysAlloc    gom.Gauge
	metricHeapTotal   gom.Gauge
)

func init() {
	startTime = time.Now()
	std = gom.NewRegistry()
	metricStartTime = GetOrRegisterGauge("Time.start")
	metricCurrentTime = GetOrRegisterGauge("Time.current")
	metricUptime = GetOrRegisterGauge("Time.uptime (s)")

	metricStartTime.Update(startTime.Unix())
}

// ConfigureOutputs will take the specified config and enable
// the required outputs.
func ConfigureOutputs(cfg *Config) (err error) {
	if err = NewFileOutput(cfg); err != nil {
		return
	}
	err = NewHTTPServer(cfg)

	if cfg.Runtime {
		metricGoRoutines = GetOrRegisterGauge("runtime.threads")
		metricGC = GetOrRegisterGauge("runtime.mem.gc")
		metricHeapObjects = GetOrRegisterGauge("runtime.mem.heap.objects")
		metricHeapAlloc = GetOrRegisterGauge("runtime.mem.heap.alloc (MB)")
		metricHeapInUse = GetOrRegisterGauge("runtime.mem.heap.inuse (MB)")
		metricHeapTotal = GetOrRegisterGauge("runtime.mem.heap.total (MB)")
		metricStackInUse = GetOrRegisterGauge("runtime.mem.stack.inuse (MB)")
		metricSysAlloc = GetOrRegisterGauge("runtime.mem.sys (MB)")
	}

	var mem runtime.MemStats
	go func() {
		for range time.Tick(time.Second) {
			metricCurrentTime.Update(time.Now().Unix())
			metricUptime.Update(int64(time.Since(startTime).Seconds()))

			runtime.ReadMemStats(&mem)
			metricGoRoutines.Update(int64(runtime.NumGoroutine()))
			metricGC.Update(int64(mem.NumGC))
			metricHeapObjects.Update(int64(mem.HeapObjects))
			metricHeapAlloc.Update(conv.BtoMB(int64(mem.Alloc)))
			metricHeapInUse.Update(conv.BtoMB(int64(mem.HeapInuse)))
			metricHeapTotal.Update(conv.BtoMB(int64(mem.TotalAlloc)))
			metricStackInUse.Update(conv.BtoMB(int64(mem.StackInuse)))
			metricSysAlloc.Update(conv.BtoMB(int64(mem.Sys)))
		}
	}()

	return
}

// NewFileOutput configures a log to periodically write metrics to.
func NewFileOutput(cfg *Config) (err error) {
	var log *out.Log
	if !cfg.FileOutput.Enabled {
		return
	}
	if log, err = out.New(cfg.FileOutput); err != nil {
		return
	}
	go func() {
		gom.WriteJSON(std, time.Minute, log)
	}()
	return
}

// NewHTTPServer sets up an HTTP server to export metrics to (if enabled)
// Metrics will be hosted at http://localhost:<port>/debug/metrics,
// where <port> is the configured port number.
func NewHTTPServer(cfg *Config) (err error) {
	if !cfg.HTTPOutput.Enabled {
		return
	}
	Exp()
	sock, err := net.Listen(
		"tcp",
		"localhost:"+strconv.Itoa(cfg.HTTPOutput.Port))
	if err != nil {
		return
	}
	go func() {
		err = http.Serve(sock, nil)
		if err != nil {
			return
		}
		out.Printf(
			"Metrics available at http://localhost:%d/debug/metrics",
			cfg.HTTPOutput.Port)
	}()
	return
}

// Exp will enable the metrics to be exported
func Exp() {
	exp.Exp(std)
}
