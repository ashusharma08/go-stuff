package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workerpool := 5
	var wg sync.WaitGroup

	dataChan := make(chan int)
	for i := 0; i < workerpool; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range dataChan {
				doTask(val)
			}
		}()
	}
	x := 0
	for range 20 {
		dataChan <- x
		x++
	}
	close(dataChan)
	wg.Wait()
}

func doTask(val int) {
	time.Sleep(2 * time.Second)
	fmt.Println("downloaded ", val)
}
