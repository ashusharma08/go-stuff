package main

type Node[K comparable] struct {
	key   K
	value interface{}
	next  *Node[K]
	prev  *Node[K]
}

func (l *LRUCache[K]) moveToTail(n *Node[K]) {
	l.remove(n)
	l.addToTail(n)

}

func (l *LRUCache[K]) remove(n *Node[K]) {
	// head<->X<->Y<->Z<->tail
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LRUCache[K]) addToTail(n *Node[K]) {
	prev := l.tail.prev

	prev.next = n
	n.prev = prev
	n.next = l.tail
	l.tail.prev = n
}

func (l *LRUCache[K]) removeFromHead() {

}

type LRUCache[K comparable] struct {
	data     map[K]*Node[K]
	size     int
	capacity int
	head     *Node[K]
	tail     *Node[K]
}

func NewLRUCache[K comparable](capacity int) *LRUCache[K] {
	head := &Node[K]{}
	tail := &Node[K]{}
	head.next = tail
	head.prev = nil
	tail.prev = head
	tail.next = nil
	return &LRUCache[K]{
		data:     make(map[K]*Node[K], capacity),
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache[K]) Get(key K) (interface{}, bool) {
	if v, ok := l.data[key]; ok {
		//item exists.. move the item and return
		l.moveToTail(v)
		return v.value, ok
	}
	return nil, false
}
func (l *LRUCache[K]) Set(key K, value interface{}) {
	if v, ok := l.data[key]; ok {
		v.value = value
		l.moveToTail(v)
		return
	}

	if l.size == l.capacity {
		n := l.head.next
		delete(l.data, n.key)
		l.removeFromHead()
	}
	node := &Node[K]{
		key:   key,
		value: value,
	}
	l.data[key] = node
	l.addToTail(node)
}
func (l *LRUCache[K]) Delete(key K) {
	if v, ok := l.data[key]; ok {
		_ = v
	}
}

type SNode struct {
	V    interface{}
	next *SNode
}

//a -> B -> c -> d -> e -> null

// prev, curr
// prev = a
// curr = B
// next = C

// p.next = null
// c.next = p
// null<- a <- B  c>d>e>null

// prev = B
// curre = C
// next =d

func main() {

}

func InPlaceRev(n *SNode) *SNode {
	var prev *SNode = nil
	var next *SNode = nil
	curr := n

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev

}
