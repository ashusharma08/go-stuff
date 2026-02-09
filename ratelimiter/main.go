package main

import (
	"log"
	"net/http"
	"time"

	"github.com/esoptra/go-prac/ratelimiter/ratelimiter"
	"github.com/esoptra/go-prac/ratelimiter/server"
)

func main() {
	if err := mainerr(); err != nil {
		log.Fatalf("error staring server")
	}
}

func RateLimiterMiddleware(rl ratelimiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rl.Limit(r.Header.Get("userid")) {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func mainerr() error {
	// rlm := ratelimiter.NewTokenBucket()
	// rlm.Init(10, 1*time.Second)
	sw := &ratelimiter.SlidingWindow{}
	sw.Init(10, 1*time.Second)

	s := &server.Server{}
	if err := http.ListenAndServe(":8000", RateLimiterMiddleware(sw)(s)); err != nil {
		return err
	}
	return nil
}
