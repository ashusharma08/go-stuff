package ratelimiter

import (
	"sync"
	"time"
)

type userBucket struct {
	tokens     int32
	lastUpdate time.Time
}
type bucket struct {
	token map[string]*userBucket
	mu    sync.RWMutex
}

type tokenBucket struct {
	interval   time.Duration
	bucketSize int32
	bucket     *bucket
}

type Options func(*tokenBucket)

func WithInterval(dur time.Duration) Options {
	return func(t *tokenBucket) {
		t.interval = dur
	}
}

func WithBucketSize(sz int32) Options {
	return func(tb *tokenBucket) {
		tb.bucketSize = sz
	}
}

func WithDefaultOptions() Options {
	return func(tb *tokenBucket) {
		tb.bucketSize = 10
		tb.interval = 1 * time.Second
	}
}

func NewTokenBucketLimiter(opts ...Options) *tokenBucket {
	t := &tokenBucket{}
	for _, o := range opts {
		o(t)
	}
	t.bucket = &bucket{
		token: make(map[string]*userBucket),
	}
	return t
}

func (t *tokenBucket) Validate(uid string) bool {
	t.bucket.mu.Lock()
	defer t.bucket.mu.Unlock()

	now := time.Now()
	ub, exists := t.bucket.token[uid]

	if !exists {
		t.bucket.token[uid] = &userBucket{
			tokens:     t.bucketSize - 1,
			lastUpdate: now,
		}
		return true
	}

	elapsed := now.Sub(ub.lastUpdate)
	refill := int32(elapsed / t.interval)

	if refill > 0 {
		ub.tokens = min(t.bucketSize, ub.tokens+refill)
		ub.lastUpdate = now
	}

	if ub.tokens > 0 {
		ub.tokens--
		return true
	}

	return false
}
