package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type Subscriber struct {
	Name         string
	Data         chan any
	P            *PubSub
	lastActivity time.Time
}

type PubSub struct {
	subscribers map[string]*Subscriber
	mu          sync.RWMutex
	lastMessage time.Time
}

func NewPubSub() *PubSub {
	p := &PubSub{
		subscribers: make(map[string]*Subscriber),
	}
	go p.monitorSubs()
	return p
}

func (p *PubSub) monitorSubs() {
	for {
		p.mu.Lock()
		for k, v := range p.subscribers {
			if time.Now().Sub(v.lastActivity) > 5*time.Second {
				fmt.Println("removing ", k)
				close(v.Data)
				delete(p.subscribers, k)
			}
		}
		p.mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func (p *PubSub) Subscribe(id string) *Subscriber {
	p.mu.Lock()
	defer p.mu.Unlock()
	subs := &Subscriber{
		Name:         id,
		Data:         make(chan any),
		lastActivity: time.Now(),
		P:            p,
	}
	p.subscribers[id] = subs
	return subs
}

func (s *Subscriber) Unsubscribe() {
	s.P.mu.Lock()
	defer s.P.mu.Unlock()
	close(s.Data)
	delete(s.P.subscribers, s.Name)
}

func (p *PubSub) Publish(message any) {
	p.lastMessage = time.Now()
	for _, v := range p.subscribers {
		go func(vv *Subscriber) {
			vv.Data <- message
			p.mu.Lock()
			vv.lastActivity = time.Now()
			p.mu.Unlock()

		}(v)
	}
}
