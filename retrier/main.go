package main

import (
	"context"
	"time"
)

func main() {

}

type retryConfig struct {
	maxRetries  int
	baseDelay   int
	ShouldRetry func(error) bool
}

func retry(ctx context.Context, cfg retryConfig, fn func() error) error {
	var err error
	for retryCount := 0; retryCount <= cfg.maxRetries; retryCount++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err = fn()
		if err == nil {
			return err
		}
		if !cfg.ShouldRetry(err) {
			return err
		}

		timer := time.NewTimer(time.Duration(cfg.baseDelay * (1 << retryCount)))

		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
	return err
}
