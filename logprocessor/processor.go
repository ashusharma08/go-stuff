package main

import (
	"context"
	"strings"
)

type processor struct {
	inChan     chan string
	dispatcher *Dispatcher
}

func newProcessor(inChan chan string, disp *Dispatcher) *processor {
	return &processor{
		dispatcher: disp,
		inChan:     inChan,
	}
}
func (p *processor) process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case l := <-p.inChan:
			p.dispatcher.broadcast(p.getParsedLog(l))
		}
	}
}

func (p *processor) getParsedLog(l string) *LogEntry {
	le := &LogEntry{}
	timestampIndex := strings.Index(l, "time")
	levelIndex := strings.Index(l, "level")
	msgIndex := strings.Index(l, "msg")
	if timestampIndex != -1 {
		startIndex := timestampIndex + len(`time="`)
		endIndex := startIndex + 25
		le.timestamp = l[startIndex:endIndex]
	}
	if levelIndex != -1 {
		endIndex := strings.Index(l[levelIndex:], " ")
		le.level = l[levelIndex+6 : levelIndex+endIndex]
	}
	if msgIndex != -1 {
		le.message = l[msgIndex+4:]
	}
	return le
}
