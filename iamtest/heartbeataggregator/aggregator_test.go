package heartbeataggregator

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_pushtobucket(b *testing.B) {
	ag := NewAggregatorToBucket(b.Context())

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			ag.PushToBucket(&Heartbeat{UserID: fmt.Sprintf("userid_%d", i), PlatformID: fmt.Sprintf("plat-%d", i), TitleID: fmt.Sprintf("title-%d", i), Timestamp: time.Now().Unix()})
			i++
		}
	})
}

func Benchmark_singleChan(b *testing.B) {
	ch := make(chan *Heartbeat, 10000)
	ag := NewAggregator()
	go ag.Process(b.Context(), ch)

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {
			ch <- &Heartbeat{UserID: fmt.Sprintf("userid_%d", i), PlatformID: fmt.Sprintf("plat-%d", i), TitleID: fmt.Sprintf("title-%d", i), Timestamp: time.Now().Unix()}
			i++
		}
	})
}

func Benchmark_shardedChan(b *testing.B) {
	ch := make([]chan *Heartbeat, 10)
	for i := range ch {
		ch[i] = make(chan *Heartbeat, 10000)
	}
	hasher := &Hasher{}
	ag := NewAggregator()
	go ag.ProcessBatch(b.Context(), ch)

	b.RunParallel(func(p *testing.PB) {
		i := 0
		for p.Next() {

			uid := fmt.Sprintf("userid_%d", i)
			idx := int(hasher.hash(uid) % uint64(10))
			ch[idx] <- &Heartbeat{UserID: uid, PlatformID: fmt.Sprintf("plat-%d", i), TitleID: fmt.Sprintf("title-%d", i), Timestamp: time.Now().Unix()}
			i++
		}
	})
}
