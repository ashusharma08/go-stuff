package ratelimiter

import (
	"sync"
	"time"
)

type SlidingWindowBucket struct {
	b  map[string][]time.Time
	mu sync.Mutex
}

type SlidingWindow struct {
	bucket       *SlidingWindowBucket
	requestLimit int
	interval     time.Duration
}

func (r *SlidingWindow) Init(limit int, interval time.Duration) {
	r.requestLimit = limit
	r.interval = interval
	r.bucket = &SlidingWindowBucket{
		b: make(map[string][]time.Time),
	}
}
func (r *SlidingWindow) Limit(userID string) bool {
	now := time.Now()
	threshold := time.Now().Add(-r.interval)
	r.bucket.mu.Lock()
	defer r.bucket.mu.Unlock()

	if _, ok := r.bucket.b[userID]; !ok {

		r.bucket.b[userID] = []time.Time{now}
		return false
	}
	for len(r.bucket.b[userID]) > 0 && r.bucket.b[userID][0].Before(threshold) {
		r.bucket.b[userID] = r.bucket.b[userID][1:]
	}
	if len(r.bucket.b[userID]) < r.requestLimit {
		r.bucket.b[userID] = append(r.bucket.b[userID], now)
		return false
	}
	return true
}
