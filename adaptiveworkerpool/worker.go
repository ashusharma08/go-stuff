package adaptiveworkerpool

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

type workerPool struct {
	config *PoolConfig

	totalWorkers  int32
	activeWorkers int32
	idleWorkers   int32

	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	ticker             time.Ticker
	controllerShutdown chan struct{}

	jobs chan Job
}

func NewWorkerPool(opts ...ConfigOptions) *workerPool {
	cfg := new(PoolConfig)
	for _, o := range opts {
		o(cfg)
	}
	ctx, cancel := context.WithCancel(context.Background())
	w := &workerPool{
		config: cfg,
		ctx:    ctx,
		cancel: cancel,
	}
	for range cfg.MinWorkers {
		w.addWorker()
	}
	//configure scalar

	return w
}

func (w *workerPool) Submit(j Job) error {
	select {
	case w.jobs <- j:
		return nil
	case <-w.ctx.Done():
		return errors.New("shutting down")
	}
}

func (w *workerPool) SubmitWithContext(ctx context.Context, j Job) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-w.ctx.Done():
		slog.Log(ctx, slog.LevelInfo, "pool is shutting down")
		return nil
	case w.jobs <- j:
		return nil
	}
}

func (w *workerPool) ShutDown() error {
	close(w.jobs)
	done := make(chan struct{})
	go func() {
		w.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-time.After(20 * time.Second):
		return fmt.Errorf("shutdown timeout exceeded")
	}
}

func (w *workerPool) addWorker() {
	atomic.AddInt32(&w.totalWorkers, 1)
	atomic.AddInt32(&w.idleWorkers, 1)
	w.wg.Add(1)
	go w.work()
}

func (w *workerPool) work() {
	defer func() {
		atomic.AddInt32(&w.totalWorkers, -1)
		atomic.AddInt32(&w.idleWorkers, -1)
		w.wg.Done()
	}()
	for {
		select {
		case j, ok := <-w.jobs:
			if !ok {
				return
			}
			atomic.AddInt32(&w.idleWorkers, -1)
			atomic.AddInt32(&w.activeWorkers, 1)

			err := j.Execute(w.ctx)
			if err != nil {
				//?? what to do
			}
			// Mark as idle
			atomic.AddInt32(&w.activeWorkers, -1)
			atomic.AddInt32(&w.idleWorkers, 1)
		case <-w.ctx.Done():
			return
		}
	}
}

// input size.
// work to be done.

type Job interface {
	Execute(ctx context.Context) error
}
