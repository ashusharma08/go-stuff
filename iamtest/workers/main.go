package workers

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Asset string
}

type Result struct {
	JobID  int
	Status string
}

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case j, ok := <-jobs:
			if !ok {
				fmt.Println("channel closed. returning")
				return nil
			}
			// Simulate "Staff" level logic: context-aware enrichment
			fmt.Printf("Worker %d processing asset %s\n", id, j.Asset)
			time.Sleep(time.Millisecond * 500) // Simulated API call
			results <- Result{JobID: j.ID, Status: "Success"}
		}
	}
}

const WORKER_POOL = 20 //get from configuration.

func main() {
	data := make([]Job, 1000)

	for idx := range data {
		data[idx] = Job{
			ID:    idx,
			Asset: fmt.Sprintf("Asset_%s", idx),
		}
	}

	var wg sync.WaitGroup
	josn := make(chan Job, 5)
	results := make(chan Result, 5)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	guard := make(chan struct{})
	go func() {
		for item := range results {
			fmt.Println("we got ", item)
		}
		guard <- struct{}{}
	}()

	wid := 0
	for range WORKER_POOL {
		wg.Add(1)
		go func(wid int) {
			err := worker(ctx, wid, josn, results, &wg)
			if err != nil {
				cancel()
			}
		}(wid)
		wid++
	}

	for _, item := range data {
		josn <- item
	}
	close(josn)
	wg.Wait()
	close(results)
	<-guard
}
