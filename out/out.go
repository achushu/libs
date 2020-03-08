package out

// TODO: Implement slice batches

import (
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/achushu/libs/concurrent"
	fs "github.com/achushu/libs/filesystem"
)

// Log is a configurable logger and file writer designed for write heavy apps.
// Can be configured to periodically rotate and to write asynchronously.
// Implements the io.Writer interface
type Log struct {
	Filename  string
	Async     bool
	Threshold PriorityLevel

	writer     io.Writer
	rotater    *lumberjack.Logger
	file       *os.File
	writeQueue *concurrent.StringSlice
	writeDelay time.Duration
	flush      chan struct{}
	stop       bool
	finished   chan struct{}
}

// Format produces a message with timestamp and priority level
func Format(pl PriorityLevel, msg string) string {
	now := time.Now().Format(time.RFC3339)
	s := now + " [" + pl.String() + "] " + msg
	return s
}

// New creates a new Log
func New(config *Config) (*Log, error) {
	var (
		writer  io.Writer
		stdFile *os.File
		rotater *lumberjack.Logger
		err     error
	)
	// default write to console if no name is specified
	writer = os.Stdout
	if config.Filename != "" {
		if config.Rotate != nil && config.Rotate.Enabled {
			rotater = &lumberjack.Logger{
				Filename:   config.Filename,
				MaxBackups: config.Rotate.MaxCount,
				MaxSize:    config.Rotate.MaxSize,
				MaxAge:     config.Rotate.MaxAge,
				Compress:   config.Rotate.Compress,
			}
			if fs.FileExists(config.Filename) && config.Rotate.RotateExisting {
				if err = rotater.Rotate(); err != nil {
					Errorf("[out/log] error rotating file: %s", err)
				}
			}
			writer = rotater
		} else {
			if stdFile, err = os.OpenFile(config.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err != nil {
				return nil, err
			}
			writer = stdFile
		}
	}

	l := &Log{
		Filename:  config.Filename,
		Async:     config.Async,
		Threshold: config.Threshold,

		writer:     writer,
		rotater:    rotater,
		file:       stdFile,
		writeQueue: concurrent.NewStringSlice(1000),
		writeDelay: time.Millisecond * 100,
		flush:      make(chan struct{}),
		stop:       false,
		finished:   make(chan struct{}),
	}
	go l.worker()
	return l, nil
}

// Logf writes a formatted string with the given priority level
// and timestamped with the current time.
// The message is only written out if PriorityLevel is
// greater than or equal to the Threshold.
func (l *Log) Logf(pl PriorityLevel, format string, args ...interface{}) {
	if pl < l.Threshold {
		return
	}
	msg := fmt.Sprintf(format, args...)
	line := Format(pl, msg)
	if l.Async {
		l.batchLine(line)
	} else {
		if _, err := l.writer.Write([]byte(line)); err != nil {
			Errorf("[out/log] error writing to file: %s", err)
		}
	}
}

// Logln writes a message with the given priority level
// and timestamped with the current time.
// The message is only written out if PriorityLevel is
// greater than or equal to the Threshold.
func (l *Log) Logln(pl PriorityLevel, args ...interface{}) {
	if pl < l.Threshold {
		return
	}
	msg := fmt.Sprintln(args...)
	line := Format(pl, msg)
	if l.Async {
		l.batchLine(line)
	} else {
		if _, err := l.writer.Write([]byte(line)); err != nil {
			Errorf("[out/log] error writing to file: %s", err)
		}
	}
}

// Write a message without additional formatting.
func (l *Log) Write(msg []byte) (int, error) {
	if l.Async {
		l.batchLine(string(msg))
		return len(msg), nil
	}
	return l.writer.Write(msg)
}

// WriteString is equivalent to Write([]byte(msg)).
func (l *Log) WriteString(msg string) (int, error) {
	return l.Write([]byte(msg))
}

// Flush triggers the log buffers to flush the output
func (l *Log) Flush() {
	l.flush <- struct{}{}
}

// Close will cleanly finish tasks and close the file
func (l *Log) Close() {
	if !l.stop {
		l.stop = true
		// Wait for writes to finish
		<-l.finished
		if l.file != nil {
			l.file.Close()
		}
		if l.rotater != nil {
			l.rotater.Close()
		}
	}
}

func (l *Log) batchLine(line string) {
	if success := l.writeQueue.Add(line); !success {
		l.Flush()
		// Try until successful
		go func() {
			for !success {
				success = l.writeQueue.Add(line)
			}
		}()
	}
}

func (l *Log) write() {
	for _, msg := range l.writeQueue.Read(true) {
		if _, err := l.writer.Write([]byte(msg)); err != nil {
			if l.Filename == "" {
				fmt.Println(Format(PriorityError, fmt.Sprintf("[out/log] error writing to stdout: %s\n", err)))
			} else {
				Errorf("[out/log] error writing to file: %s", err)
			}
		}
	}
	l.writeQueue.ReadDone()
}

func (l *Log) worker() {
	var timer = time.NewTimer(l.writeDelay)
	for !l.stop {
		// Wait for the write delay or a call to flush the buffers
		timer.Reset(l.writeDelay)
		select {
		case <-l.flush:
			timer.Stop()
		case <-timer.C:
		}
		l.write()
	}
	l.write()
	// Signal completion
	l.finished <- struct{}{}

}
