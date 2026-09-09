package main

import (
	"hash/maphash"
	"sync"
)

type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

type LRUShard[K comparable, V any] struct {
	shards []*LRUCache[K, V]
	hasher *Hasher[K]
}

type LRUCache[K comparable, V any] struct {
	data     map[K]*node[K, V]
	capacity int
	size     int
	mu       sync.RWMutex
	head     *node[K, V]
	tail     *node[K, V]
}

func NewCache[K comparable, V any](numShard, capShard int) *LRUShard[K, V] {
	l := &LRUShard[K, V]{
		shards: make([]*LRUCache[K, V], numShard),
		hasher: NewHasher[K](),
	}
	for i := range l.shards {
		l.shards[i] = &LRUCache[K, V]{
			capacity: capShard,
			data:     make(map[K]*node[K, V]),
			head:     nil,
			tail:     nil,
		}
	}
	return l
}

func (l *LRUShard[K, V]) getShard(key K) int {
	return int(l.hasher.Hash(key)) % len(l.shards)
}

func (ls *LRUShard[K, V]) Set(k K, v V) {
	l := ls.shards[ls.getShard(k)]

	l.mu.Lock()
	defer l.mu.Unlock()
	n := &node[K, V]{
		key:   k,
		value: v,
	}
	l.data[k] = n
	l.tail = n
}

func (l *LRUCache[K, V]) AddToFront(n *node[K, V]) {
	//its already empty
	if l.size == 0 {
		l.head = n
		l.tail = n
		return
	}
	currentHead := l.head
	currentHead.next = n
	n.prev = currentHead
	l.head = n
}

func (l *LRUCache[K, V]) MoveToFront(n *node[K, V]) {
	prev := n.prev
	next := n.next

	if l.head == n {
		//already at front
		return
	}
	prev.next = next
	next.prev = prev

	n.prev = l.head
	l.head.next = n
	n.next = nil
	l.head = n
}

func (l *LRUCache[K, V]) RemoveFromTail() {
	n := l.tail

	l.tail = n.next
	n.next.prev = l.tail
	n.next = nil
}

func (ls *LRUShard[K, V]) Get(k K) (v any, ok bool) {
	l := ls.shards[ls.getShard(k)]
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, exists := l.data[k]
	n := v.(*node[K, V])

	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
	l.tail.next = n
	n.next = nil
	n.prev = l.tail

	l.tail = n
	return n, exists
}

type Hasher[K comparable] struct {
	seed maphash.Seed
}

func NewHasher[K comparable]() *Hasher[K] {
	return &Hasher[K]{
		seed: maphash.MakeSeed(),
	}
}

func (h *Hasher[K]) Hash(k K) uint64 {
	return maphash.Comparable(h.seed, k)
}
