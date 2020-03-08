package metrics

import gom "github.com/rcrowley/go-metrics"

var (
	std gom.Registry
)

func GetOrRegister(name string, i interface{}) interface{} {
	return std.GetOrRegister(name, i)
}

func GetOrRegisterCounter(name string) gom.Counter {
	return gom.GetOrRegisterCounter(name, std)
}

func GetOrRegisterGauge(name string) gom.Gauge {
	return gom.GetOrRegisterGauge(name, std)
}

func GetOrRegisterHistogram(name string, sample gom.Sample) gom.Histogram {
	return gom.GetOrRegisterHistogram(name, std, sample)
}

func GetOrRegisterMeter(name string) gom.Meter {
	return gom.GetOrRegisterMeter(name, std)
}

func GetOrRegisterTimer(name string) gom.Timer {
	return gom.GetOrRegisterTimer(name, std)
}

func Register(name string, i interface{}) error {
	return std.Register(name, i)
}

func Unregister(name string) {
	std.Unregister(name)
}

func UnregisterAll() {
	std.UnregisterAll()
}
