package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

// 6. Fan-in Logger (Backpressure Handling)
// Build logging system:
// Many goroutines produce logs
// One writer goroutine writes to file
// Buffered channel
// Handle channel full without blocking system
// Follow-ups
// Drop logs vs spill to disk
// Batching
// Graceful shutdown flush
func main() {

}

type Level int

const (
	TRACE Level = 1
	DEBUG Level = 2
	INFO  Level = 3
	WARN  Level = 4
	ERROR Level = 5
)

func (l Level) String() string {
	switch l {
	case TRACE:
		return "trace"
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "invalid"
	}
}

type Logger struct {
	Level       Level // this defines the minimum level
	ch          chan log
	destination io.Writer
	ctx         context.Context
	cancel      context.CancelFunc
}

type log struct {
	level     Level
	timestamp time.Time
	message   string
}

func (l log) write() string {
	return fmt.Sprintf("%s %s %s", l.timestamp.GoString(), "level="+l.level.String(), l.message)
}

type Options func(*Logger)

func WithDefaults() Options {
	return func(l *Logger) {
		l.Level = TRACE
	}
}

func NewLogger(opts ...Options) (*Logger, error) {
	ctx, cancel := context.WithCancel(context.Background())
	l := &Logger{
		ch: make(chan log, 100),
	}
	if len(opts) == 0 {
		opts = append(opts, WithDefaults())
	}
	for _, o := range opts {
		o(l)
	}
	l.ctx = ctx
	l.cancel = cancel
	f, err := os.Create("logfile" + fmt.Sprint(time.Now().Unix()))
	if err != nil {
		return nil, fmt.Errorf("error initializing logger")
	}
	l.destination = f
	go l.log()
	return l, nil
}

func (l *Logger) close() {
	l.cancel()
	close(l.ch)
}

func (l *Logger) Tracef(msg string, args ...any) {
	l.ch <- log{
		level:     TRACE,
		timestamp: time.Now().UTC(),
		message:   fmt.Sprintf(msg, args...),
	}
}
func (l *Logger) Debugf(msg string, args ...any) {
	l.ch <- log{
		level:     DEBUG,
		timestamp: time.Now().UTC(),
		message:   fmt.Sprintf(msg, args...),
	}
}
func (l *Logger) Infof(msg string, args ...any) {
	l.ch <- log{
		level:     INFO,
		timestamp: time.Now().UTC(),
		message:   fmt.Sprintf(msg, args...),
	}
}
func (l *Logger) Warnf(msg string, args ...any) {
	l.ch <- log{
		level:     WARN,
		timestamp: time.Now().UTC(),
		message:   fmt.Sprintf(msg, args...),
	}
}
func (l *Logger) Errorf(msg string, args ...any) {
	l.ch <- log{
		level:     ERROR,
		timestamp: time.Now().UTC(),
		message:   fmt.Sprintf(msg, args...),
	}
}
func (l *Logger) log() {
	for lg := range l.ch { // drains until channel closed
		l.destination.Write([]byte(lg.write()))
	}

	// after drain → safe to close file
	if wc, ok := l.destination.(io.WriteCloser); ok {
		wc.Close()
	}
}
