package lrucache

import (
	"hash/maphash"
	"sync"
)

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
	prev  *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	data       map[K]*Node[K, V]
	size       int
	head, tail *Node[K, V]

	mu sync.RWMutex
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

type ShardedLRU[K comparable, V any] struct {
	shards []*LRUCache[K, V]
	hasher *Hasher[K]
}

func NewShardedLRU[K comparable, V any](numberOfShards, shardCapacity int) *ShardedLRU[K, V] {
	l := &ShardedLRU[K, V]{
		shards: make([]*LRUCache[K, V], numberOfShards),
		hasher: NewHasher[K](),
	}

	for i := range l.shards {
		head, tail := &Node[K, V]{}, &Node[K, V]{}
		head.next = tail
		tail.prev = head
		l.shards[i] = &LRUCache[K, V]{
			size: shardCapacity,
			data: make(map[K]*Node[K, V]),
			head: head,
			tail: tail,
		}
	}
	return l
}

// Tail is where we keep the most recently used item.
// Head is where we keep the least recently used item
func NewLRUCache[K comparable, V any](size int) *LRUCache[K, V] {
	head, tail := &Node[K, V]{}, &Node[K, V]{}
	head.next = tail
	tail.prev = head

	return &LRUCache[K, V]{
		data: make(map[K]*Node[K, V]),
		size: size,
		head: head,
		tail: tail,
	}
}

func (l *ShardedLRU[K, V]) Set(key K, value V) {
	shard := l.shards[l.getShard(key)]

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if node, ok := shard.data[key]; ok {
		node.value = value
		shard.moveToTail(node)
		return
	}

	newNode := &Node[K, V]{key: key, value: value}
	shard.data[key] = newNode
	shard.addToTail(newNode)

	if len(shard.data) > shard.size {
		// Evict the least recently used (the one after head)
		shard.remove(shard.head.next)
	}
}

func (l *ShardedLRU[K, V]) getShard(key K) int {
	return int(l.hasher.Hash(key) % uint64(len(l.shards)))
}

// Get a value from Cache
func (l *ShardedLRU[K, V]) Get(key K) (V, bool) {
	shard := l.shards[l.getShard(key)]

	shard.mu.Lock()
	defer shard.mu.Unlock()

	node, ok := shard.data[key]
	if !ok {
		var zero V
		return zero, false
	}

	shard.moveToTail(node)
	return node.value, true
}

func (l *LRUCache[K, V]) addToTail(n *Node[K, V]) {
	prev := l.tail.prev

	prev.next = n
	n.prev = prev
	n.next = l.tail
	l.tail.prev = n
}

func (l *LRUCache[K, V]) moveToTail(n *Node[K, V]) {
	l.remove(n)
	l.addToTail(n)
}

func (l *LRUCache[K, V]) evict(n *Node[K, V]) {
	l.remove(n)
	delete(l.data, n.key)
}

func (l *LRUCache[K, V]) remove(n *Node[K, V]) {
	prev := n.prev
	next := n.next

	prev.next = next
	next.prev = prev
}
