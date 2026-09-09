package lrucache

import (
	"fmt"
	"testing"
)

func Benchmark_LRUCache(b *testing.B) {
	capacity := 10000
	size := 16

	cache := NewShardedLRU[string, int](size, capacity)

	i := 0
	for range capacity {
		cache.Set(fmt.Sprintf("key-%d", i), i)
		i++
	}

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Get(fmt.Sprintf("key-%d", i%capacity))
			i++
		}
	})
}

func Benchmark_UnshardedLRUCache(b *testing.B) {
	capacity := 160000
	cache := NewUnshardedLRUCache[string, int](capacity)
	i := 0
	for range 160000 {
		cache.Set(fmt.Sprintf("key-%d", i), i)
		i++
	}
	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			cache.Get(fmt.Sprintf("key-%d", i%capacity))
			i++
		}
	})
}
