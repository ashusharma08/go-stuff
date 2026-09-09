package main

type LogLevel string

const (
	Info    LogLevel = "info"
	Debug   LogLevel = "debug"
	Trace   LogLevel = "trace"
	Error   LogLevel = "error"
	Unknown LogLevel = "unknown"
)

type LogEntry struct {
	level     string
	timestamp string
	message   string
	labels    map[string]any
}

type Dispatcher struct {
	subscribers []chan *LogEntry
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) register(ch chan *LogEntry) {
	d.subscribers = append(d.subscribers, ch)
}
func (d *Dispatcher) broadcast(entry *LogEntry) {
	for _, c := range d.subscribers {
		c <- entry
	}
}
