package main

import "sync"

// Bounded Blocking Queue
// Implement:
// 		Fixed capacity queue
// 		Push() blocks when full
// 		Pop() blocks when empty
// 		Multiple producers + consumers
// 		No busy loop
// Follow-ups
// 		Add timeout
// 		Add non-blocking TryPush/TryPop
// 		Measure contention

func main() {

}

type Queue struct {
	q        []interface{}
	head     int
	tail     int
	size     int
	capacity int
	mu       sync.Mutex
	notfull  *sync.Cond
	notempty *sync.Cond
}

func (q *Queue) Push(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Wait until space available
	for q.size == q.capacity {
		q.notfull.Wait()
	}

	q.q[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	q.size++

	// Wake one waiting consumer
	q.notempty.Signal()

}

func (q *Queue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	for q.size == 0 {
		q.notempty.Wait()
	}

	v := q.q[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--

	q.notfull.Signal()

	return v
}

func NewQueue(capacity int) *Queue {
	q := &Queue{
		q:        make([]interface{}, 0, capacity),
		capacity: capacity,
	}
	q.notfull = sync.NewCond(&q.mu)
	q.notempty = sync.NewCond(&q.mu)
	return q
}
