package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(":4222")
	if err != nil {
		log.Fatalf("fatal error %#v", err)
	}

	nc.QueueSubscribe("TEST.SUBJECT", "my.queue", func(msg *nats.Msg) {

		fmt.Println(string(msg.Data))
	})

	time.Sleep(30 * time.Second)
}
