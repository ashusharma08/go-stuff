package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(":4222")
	if err != nil {
		log.Fatalf("fatal error %#v", err)
	}

	nc.Publish("TEST.SUBJECT", []byte("this is a message"))
	time.Sleep(30 * time.Second)
}
