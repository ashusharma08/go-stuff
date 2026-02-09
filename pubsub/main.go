package main

import (
	"fmt"
	"time"

	"github.com/esoptra/go-prac/pubsub/pubsub"
)

func main() {
	ps := pubsub.NewPubSub()

	sub1 := ps.Subscribe("Ashish")
	sub2 := ps.Subscribe("Anjna")

	go func(s *pubsub.Subscriber) {
		for val := range s.Data {
			fmt.Println("received data for sub1: ", val)
		}
	}(sub1)
	go func(s *pubsub.Subscriber) {
		for val := range s.Data {
			fmt.Println("received data for sub2: ", val)
		}
	}(sub2)

	ps.Publish("Hellow from the publisher")
	ps.Publish("Hellow from the publisher for the 2nd time")
	time.Sleep(2 * time.Second)
	sub1.Unsubscribe()

	ps.Publish("new message")
	ps.Publish("second enw message")
	time.Sleep(2 * time.Second)

	sub2.Unsubscribe()
}
