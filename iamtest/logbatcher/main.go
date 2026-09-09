package main

import (
	"context"
	"time"
)

func main() {

}

type Logger struct {
	ctx     context.Context
	cancel  context.CancelFunc
	flusher func([]string)
	bc      chan string
	bucket  []string
}

func NewLogger(ctx context.Context, flushFunc func([]string)) *Logger {
	bctx, cancel := context.WithCancel(ctx)
	b := &Logger{
		ctx:     bctx,
		cancel:  cancel,
		flusher: flushFunc,
		bc:      make(chan string, 100),
		bucket:  make([]string, 0, 100),
	}
	go b.batcher()
	return b
}

func (b *Logger) batcher() {
	var timerChan <-chan time.Time
	var timer *time.Timer

	flush := func() {
		if len(b.bucket) > 0 {
			cp := make([]string, len(b.bucket))
			copy(cp, b.bucket)
			b.flusher(cp)
			b.bucket = b.bucket[:0]
		}
		if timer != nil {
			timer.Stop()
			timer = nil
			timerChan = nil
		}
	}

	for {
		select {
		case <-b.ctx.Done():
			flush()
			return
		case <-timerChan:
			flush()
		case msg := <-b.bc:
			if len(b.bucket) == 0 {
				timer = time.NewTimer(200 * time.Millisecond)
				timerChan = timer.C
			}
			b.bucket = append(b.bucket, msg)
			if len(b.bucket) == 100 {
				flush()
			}
		}
	}
}

func (b *Logger) Log(msg string) {
	b.bc <- msg
}
