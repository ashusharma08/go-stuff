package ratelimiter

import (
	"sync"
	"time"

	"log"
)

type Limiter interface {
	Init(limit int, interval time.Duration)
	Limit(user string) bool
}

type Bucket struct {
	token map[string]int
	mu    sync.Mutex
}

type TokenBucket struct {
	b *Bucket
}

func NewTokenBucket() *TokenBucket {
	r := &TokenBucket{
		b: &Bucket{
			token: make(map[string]int),
		},
	}
	return r
}
func (r *TokenBucket) Init(limit int, interval time.Duration) {

	go func() {
		for {
			time.Sleep(1 * time.Second)
			r.b.mu.Lock()
			// Copy user IDs to avoid holding lock for long
			users := make([]string, 0, len(r.b.token))
			for userID := range r.b.token {
				users = append(users, userID)
			}
			r.b.mu.Unlock()

			// Reset tokens without holding lock
			for _, userID := range users {
				log.Printf("Resetting tokens for userID: %s", userID)
				r.b.mu.Lock()
				r.b.token[userID] = 10
				r.b.mu.Unlock()
			}
		}
	}()
}

func (r *TokenBucket) Limit(userID string) bool {
	r.b.mu.Lock()
	defer r.b.mu.Unlock()
	if x, ok := r.b.token[userID]; ok {
		r.b.token[userID]--
		log.Printf("%d", r.b.token[userID])
		if x > 0 {
			return false
		}
		return true
	}
	log.Printf("adding new user %s", userID)
	r.b.token[userID] = 10

	return false
}
