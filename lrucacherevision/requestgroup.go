package lrucache

import (
	"context"
	"sync"
)

type RequestGroup[K comparable, V any] struct {
	lruCache *ShardedLRU[K, V]
	callMap  map[K]*call[V]
	mu       sync.Mutex
}

type call[V any] struct {
	err   error
	value V
	ch    chan struct{}
}

func NewRequestGroup[K comparable, V any](cache *ShardedLRU[K, V]) *RequestGroup[K, V] {
	return &RequestGroup[K, V]{
		lruCache: NewShardedLRU[K, V](16, 100),
		callMap:  make(map[K]*call[V]),
	}
}
func (r *RequestGroup[K, V]) Do(ctx context.Context, key K, fn func(context.Context) (V, error)) (V, error) {

	var val V

	v, ok := r.lruCache.Get(key)
	if ok {
		return v, nil
	}
	r.mu.Lock()
	c, ok := r.callMap[key]
	if ok {
		r.mu.Unlock()
		select {
		case <-c.ch:
			return c.value, c.err
		case <-ctx.Done():
			return val, ctx.Err()
		}
	}

	call := &call[V]{
		ch: make(chan struct{}),
	}
	r.mu.Lock()
	r.callMap[key] = call
	r.mu.Unlock()

	go call.execute(ctx, fn)

	select {
	case <-c.ch:
		r.mu.Lock()
		delete(r.callMap, key)
		r.mu.Unlock()

		if call.err == nil {
			r.lruCache.Set(key, call.value)
		}
		return call.value, call.err
	case <-ctx.Done():
		return val, ctx.Err()
	}

}

func (c *call[V]) execute(ctx context.Context, fn func(context.Context) (V, error)) {
	res, err := fn(ctx)
	c.value = res
	c.err = err
	close(c.ch)
}
