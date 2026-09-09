package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	processor(ctx)

}

// You have 100 workers. You need to process 1 million items. However, if the error rate (not just a single error) exceeds 5% in a rolling 10-second window, you must trigger a graceful shutdown.

func worker(ctx context.Context, ch <-chan string, success, failure chan<- struct{}) {

	defer func() {
		if r := recover(); r != nil {
			//panic recover
			failure <- struct{}{}
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-ch:
			if !ok {
				return //channel closed
			}
			rv := rand.IntN(20)
			time.Sleep(time.Duration(rv * int(time.Second)))
			success <- struct{}{}
		}
	}
}

func processor(ctx context.Context) {
	pctx, cancel := context.WithCancel(ctx)
	defer cancel()
	success, failure := make(chan struct{}, 1000), make(chan struct{}, 1000)

	var totalCount int32
	var errorCount int32

	go func() {
		for {
			select {
			case <-pctx.Done():
				return
			case <-success:
				atomic.AddInt32(&totalCount, 1)
			case <-failure:
				atomic.AddInt32(&totalCount, 1)
				atomic.AddInt32(&errorCount, 1)
			}
		}
	}()

	go func() {
		type bucket struct {
			total, errors int32
		}
		bc := make([]*bucket, 10)

		var lastTotal, lastError int32
		cursor := 0

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				currTotal, currErr := atomic.LoadInt32(&totalCount), atomic.LoadInt32(&errorCount)
				deltaTotal, deltaErr := currTotal-lastTotal, currErr-lastError

				bc[cursor] = &bucket{
					total:  deltaTotal,
					errors: deltaErr,
				}
				cursor = (cursor + 1) % 10
				var winTotal, winError int32
				for _, item := range bc {
					winTotal += item.total
					winError += item.errors
				}
				if winTotal > 0 {
					rate := float64(winError) / float64(winTotal)
					if rate > 5.0 {
						cancel()
						return
					}
				}
				lastTotal, lastError = currTotal, currErr
			case <-pctx.Done():
				return
			}
		}
	}()

	ch := make(chan string)

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(pctx, ch, success, failure)
		}()
	}

	go func() {
		defer close(ch)
		i := 0
		for range 1_000_000 {
			select {
			case <-pctx.Done():
				return
			case ch <- fmt.Sprint(i):
				//default:
				//handle backpressure
			}
		}
	}()

	wg.Wait()
}
