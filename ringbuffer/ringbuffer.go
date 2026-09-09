package ringbuffer

import (
	"runtime"
	"sync/atomic"
)

type RingBuffer[T any] struct {
	buffer []T
	ready  []uint32
	mask   uint64
	head   uint64
	tail   uint64
	size   uint64
}

func (rb *RingBuffer[T]) Push(item T) {
	for {
		head := atomic.LoadUint64(&rb.head)
		tail := atomic.LoadUint64(&rb.tail)

		if head-tail >= rb.size {
			runtime.Gosched()
			continue
		}

		if atomic.CompareAndSwapUint64(&rb.head, head, head+1) {
			idx := head & rb.mask
			atomic.StoreUint32(&rb.ready[idx], 1)
			rb.buffer[idx] = item
			atomic.StoreUint32(&rb.ready[idx], 2)
			return
		}
	}
}
func (rb *RingBuffer[T]) Pop() T {
	for {
		head := atomic.LoadUint64(&rb.head)
		tail := atomic.LoadUint64(&rb.tail)

		if head == tail {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapUint64(&rb.tail, tail, tail+1) {
			idx := tail & rb.mask
			for atomic.LoadUint32(&rb.ready[idx]) != 2 {
				runtime.Gosched()
			}

			val := rb.buffer[idx]

			var zer T
			rb.buffer[idx] = zer
			atomic.StoreUint32(&rb.ready[idx], 0)
			return val
		}
	}
}
