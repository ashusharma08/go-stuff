package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	t := time.Now()
	var wg1 sync.WaitGroup
	m := &MutexCounter{}
	for range 1000 {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.i++
		}()
	}
	wg1.Wait()
	fmt.Println("time taken by mutex", time.Since(t).Seconds())

	t = time.Now()
	ac := &AtomicCounter{}
	var wg2 sync.WaitGroup
	for range 1000 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			ac.i.Add(1)
		}()
	}
	wg2.Wait()
	fmt.Println("time taken by atomic", time.Since(t).Seconds())

	t = time.Now()
	i := 0
	counterChan := make(chan struct{}, 1000)
	var recwg sync.WaitGroup
	recwg.Add(1)
	go func() {
		defer recwg.Done()
		for range counterChan {
			i++
		}
	}()
	var snwg sync.WaitGroup
	for range 1000 {
		snwg.Add(1)
		go func() {
			defer snwg.Done()
			counterChan <- struct{}{}
		}()
	}
	snwg.Wait()
	close(counterChan)
	recwg.Wait()
	fmt.Println("time taken by channel", time.Since(t).Seconds())
}

type MutexCounter struct {
	i  int
	mu sync.Mutex
}

type AtomicCounter struct {
	i atomic.Int32
}
