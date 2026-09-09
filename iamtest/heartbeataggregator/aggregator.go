package heartbeataggregator

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/cespare/xxhash/v2"
)

type Heartbeat struct {
	UserID     string
	PlatformID string
	TitleID    string
	Timestamp  int64
}

type Aggregator struct {
	outChan  chan []*bucket
	bucket   *shard
	duration time.Duration
}

type shard struct {
	mu       sync.Mutex
	capacity int
	shards   []*bucket
	hasher   *Hasher
}

type bucket struct {
	data map[string]*Heartbeat
	mu   sync.Mutex
}

func NewAggregator() *Aggregator {
	ag := &Aggregator{
		outChan:  make(chan []*bucket, 10000),
		duration: 5 * time.Second,
		bucket: &shard{
			capacity: 10,
			shards:   make([]*bucket, 10),
			hasher:   &Hasher{},
		},
	}
	for idx := range ag.bucket.shards {
		ag.bucket.shards[idx] = &bucket{
			data: make(map[string]*Heartbeat),
		}
	}

	return ag
}

func NewAggregatorToBucket(ctx context.Context) *Aggregator {
	ag := &Aggregator{
		outChan:  make(chan []*bucket, 10000),
		duration: 5 * time.Second,
		bucket: &shard{
			capacity: 10,
			shards:   make([]*bucket, 10),
			hasher:   &Hasher{},
		},
	}
	for idx := range ag.bucket.shards {
		ag.bucket.shards[idx] = &bucket{
			data: make(map[string]*Heartbeat),
		}
	}
	go ag.runTicker(ctx)
	go ag.consume(context.Background())
	return ag
}

func (a *Aggregator) runTicker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			a.flush(ctx)
			return
		case <-ticker.C:
			a.flush(ctx)
		}
	}
}

func (a *Aggregator) PushToBucket(h *Heartbeat) {
	a.pushToBucket(h)
}

var bucketPool = sync.Pool{
	New: func() any {
		return &bucket{data: make(map[string]*Heartbeat, 10000)}
	},
}

func (a *Aggregator) flushShard(ctx context.Context, workerID int) error {
	// 1. Get the current bucket the worker has been filling
	oldBucket := a.bucket.shards[workerID]

	// 2. Immediately replace it with a fresh/empty one from the Pool
	// This takes nanoseconds. The worker is now ready for the next 5s.
	a.bucket.shards[workerID] = bucketPool.Get().(*bucket)

	// 3. Send the OLD bucket to the outgoing pipeline
	select {
	case a.outChan <- []*bucket{oldBucket}:
		// Success: The consumer will handle DB writes and ReturnBucket() to the pool
	default:
		//log.Warn("Dropped batch for shard", "shard_id", workerID)
	}
	return nil
}

func (a *Aggregator) flush(ctx context.Context) error {
	newShards := make([]*bucket, a.bucket.capacity)
	for i := range newShards {
		newShards[i] = bucketPool.Get().(*bucket)
	}

	a.bucket.mu.Lock()
	oldBucket := a.bucket.shards
	a.bucket.shards = newShards
	a.bucket.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case a.outChan <- oldBucket:
	default:
		//drop something for backpressure
	}

	return nil
}

func (a *Aggregator) consume(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case shards := <-a.outChan:
			for _, b := range shards {
				//write to db
				time.Sleep(time.Duration(rand.IntN(10) * int(time.Second)))
				for k := range b.data {
					delete(b.data, k)
				}
				bucketPool.Put(b)
			}
		}
	}
}

func (a *Aggregator) ProcessBatch(ctx context.Context, dataChan []chan *Heartbeat) error {
	var wg sync.WaitGroup
	go a.consume(ctx)
	wg.Add(1)
	for i := 0; i < a.bucket.capacity; i++ {
		ch := dataChan[i]
		shard := a.bucket.shards[i]
		go func(wid int, c chan *Heartbeat, s *bucket) {
			ticker := time.NewTicker(a.duration)
			for {
				select {
				case hb := <-c:
					// Only THIS goroutine ever touches THIS shard
					s.data[hb.UserID] = hb
				case <-ticker.C:
					a.flushShard(ctx, wid)
				case <-ctx.Done():
					return
				}
			}
		}(i, ch, shard)
	}
	defer wg.Wait()
	return nil
}

func (a *Aggregator) Process(ctx context.Context, dataChan chan *Heartbeat) error {
	var wg sync.WaitGroup
	go a.consume(ctx)
	wg.Add(1)
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(a.duration)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					go a.flush(ctx)
					return
				case v, ok := <-dataChan:
					if !ok {
						go a.flush(ctx)
						return
					}
					a.pushToBucket(v)
				case <-ticker.C:
					go a.flush(ctx)
				}
			}
		}()
	}
	defer wg.Wait()
	return nil
}

func (a *Aggregator) pushToBucket(v *Heartbeat) {
	idx := int(a.bucket.hasher.hash(v.UserID) % uint64(a.bucket.capacity))
	shard := a.bucket.shards[idx]
	shard.mu.Lock()
	defer shard.mu.Unlock()
	userHeartBeat, ok := shard.data[v.UserID]
	if !ok || v.Timestamp > userHeartBeat.Timestamp {
		shard.data[v.UserID] = v
	}
}

type exporter interface {
	Export(context.Context, chan []*Heartbeat)
}

type Hasher struct {
}

func (h *Hasher) hash(key string) uint64 {
	return xxhash.Sum64String(key)
}
