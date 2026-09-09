package lrucache

// Tail is where we keep the most recently used item.
// Head is where we keep the least recently used item
func NewUnshardedLRUCache[K comparable, V any](size int) *LRUCache[K, V] {
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

// Adds new item to the cache
func (l *LRUCache[K, V]) Set(key K, value V) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// check if the key already exists.
	// If yes, the update the value and move it to the tail
	existing, ok := l.data[key]
	if ok {
		existing.value = value
		l.moveToTail(existing)
		return
	}
	// New key. So create a node and add it to the tail
	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}

	l.data[key] = newNode
	l.addToTail(newNode)

	// If we are exceeding capacity, then remove the oldest guy from head
	if len(l.data) > l.size {
		l.evict(l.head.next)
	}
}

// Get a value from Cache
func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	var res V

	existing, ok := l.data[key]
	if !ok {
		return res, false
	}
	// The item exists, so we move it to the tail
	l.moveToTail(existing)

	return existing.value, true
}
