package main

import (
	"context"
	"fmt"
	"strings"
)

type aggregator interface {
	process(context.Context, chan *LogEntry)
	print()
}

type errorAggregator struct {
	errorLogs []*LogEntry
}

func newErrorAggregator() *errorAggregator {
	return &errorAggregator{
		errorLogs: make([]*LogEntry, 0),
	}
}

func (e *errorAggregator) process(ctx context.Context, le chan *LogEntry) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-le:
			if !ok {
				return
			}
			if log.level == "error" {
				e.errorLogs = append(e.errorLogs, log)
			}
		}
	}
}

func (e *errorAggregator) print() {
	if len(e.errorLogs) > 0 {
		for _, item := range e.errorLogs {
			fmt.Println("####ERRORAGGREGATOR", item.timestamp, item.message)
		}
	} else {
		fmt.Println("####ERRORAGGREGATOR no error logs####")
	}
}

type keywordAggregator struct {
	keywords []string
	logs     []*LogEntry
}

func newKeywordAggregator(kw []string) *keywordAggregator {
	return &keywordAggregator{
		keywords: kw,
		logs:     make([]*LogEntry, 0),
	}
}

func (e *keywordAggregator) process(ctx context.Context, le chan *LogEntry) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-le:
			if !ok {
				return
			}
			for _, k := range e.keywords {
				if strings.Contains(log.message, k) {
					e.logs = append(e.logs, log)
				}
			}
		}
	}
}

func (e *keywordAggregator) print() {
	if len(e.logs) > 0 {
		for _, item := range e.logs {
			fmt.Println("####KEYWORDAGGREGATOR", item.timestamp, item.message)
		}
	} else {
		fmt.Println("####KEYWORDAGGREGATOR no error logs####")
	}
}
