package main

import (
	"context"
	"sync"

	"github.com/esoptra/go-prac/consumerproducer/consumer"
	"github.com/esoptra/go-prac/consumerproducer/producer"
)

func main() {
	prod := producer.GetProducers(4)
	cons := consumer.GetConsumers(2, prod)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{}, 2) // Channel to signal consumer completion
	var pwg sync.WaitGroup
	for _, p := range prod {
		pwg.Add(1)
		go func(pr *producer.Producer) {
			defer pwg.Done()
			pr.Produce(ctx)
		}(p)
	}
	for _, c := range cons {
		go func(cm *consumer.Consumer) {
			cm.Consume()
			done <- struct{}{} // Signal completion
		}(c)
	}

	for i := 0; i < 2; i++ {
		<-done
	}
	cancel()
	pwg.Wait()
}
