package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

// The Challenge: "The Distributed Rate-Limited Web Crawler"
// The Scenario:
// You are building a service for WBD that needs to crawl metadata from 1,000 different external partner domains (e.g., IMDB, Rotten Tomatoes, etc.).

// The Constraints:

// Global Concurrency: You have a worker pool of 50 goroutines.

// Per-Domain Rate Limit: You must not hit any single domain more than 3 times per second.

// Efficiency: If one domain is slow, your workers shouldn't sit idle; they should move on to tasks for other domains that haven't hit their rate limit.

// Graceful Exit: The system must shut down cleanly via context and report how many URLs were successfully crawled and how many were skipped due to the shutdown.

func main() {

}

type bucket struct {
	seedDomains map[string]*rate.Limiter
	urls        []string
	mu          sync.RWMutex
}

func NewBucket(seedDomains []string) *bucket {
	b := &bucket{
		seedDomains: make(map[string]*rate.Limiter),
		urls:        make([]string, 0, 1_000_000),
	}
	for _, item := range seedDomains {
		b.seedDomains[item] = rate.NewLimiter(rate.Every(1*time.Second), 3)
		b.urls = append(b.urls, item)
	}
	return b
}

func (b *bucket) process(ctx context.Context) {
	readChan := make(chan string, 10000)

	var atom atomic.Int32
	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case url, ok := <-readChan:
					if !ok {
						return
					}
					atom.Add(-1)
					domain := b.getDomain(url)
					rl, ok := b.seedDomains[domain]
					if !ok {
						continue
					}
					if !rl.Allow() {
						atom.Add(1)
						readChan <- url
						continue
					}
					res, _ := b.work(url)
					//handle error
					for _, item := range res {
						atom.Add(1)
						readChan <- item
					}

					if atom.Load() == 0 {
						close(readChan)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		for _, item := range b.urls {
			atom.Add(1)
			readChan <- item
		}
	}()
	wg.Wait()
}

func (b *bucket) work(u string) ([]string, error) {
	rn := rand.IntN(10)
	if rn < 3 {
		return nil, nil
	}
	return []string{
		fmt.Sprintf("%s/%d", u, rn),
		fmt.Sprintf("%s/%d", u, rn+1),
	}, nil
}

func (b *bucket) getDomain(inURL string) string {
	u, _ := url.Parse(inURL)
	return u.Host
}
