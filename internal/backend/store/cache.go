package store

import (
	"container/heap"
	"sync"
	"time"
)

var Cache *CacheInstance

type ItemHeap []*Item

type Item struct {
	Key        string
	Value      any
	Expiration int64
	Index      int
}

type CacheInstance struct {
	store  map[string]*Item
	queue  ItemHeap
	mu     sync.RWMutex
	signal chan struct{}
}

func init() {
	Cache = &CacheInstance{
		store:  make(map[string]*Item),
		queue:  make(ItemHeap, 0),
		signal: make(chan struct{}),
	}
	heap.Init(&Cache.queue)
	go Cache.cleanup()

	println("✓ Started cache")
}

func (c *CacheInstance) Get(key string) (value any, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.store[key]
	if !ok || time.Now().UnixNano() > item.Expiration {
		delete(c.store, key)
		return nil, false
	}
	return item.Value, true
}

func (c *CacheInstance) Set(key string, value any, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(duration).UnixNano()
	if item, ok := c.store[key]; ok {
		item.Value = value
		item.Expiration = expiration
		heap.Fix(&c.queue, item.Index)
	} else {
		item = &Item{
			Key:        key,
			Value:      value,
			Expiration: expiration,
		}
		heap.Push(&c.queue, item)
		c.store[key] = item
	}

	c.signal <- struct{}{}
}

func (c *CacheInstance) cleanup() {
	for {
		c.mu.Lock()
		for len(c.queue) > 0 {
			item := heap.Pop(&c.queue).(*Item)
			if time.Now().UnixNano() > item.Expiration {
				delete(c.store, item.Key)
			} else {
				heap.Push(&c.queue, item)
				break
			}
		}
		c.mu.Unlock()

		if len(c.queue) > 0 {
			<-time.After(time.Until(time.Unix(0, c.queue[0].Expiration)))
		} else {
			<-c.signal
		}
	}
}

func (h ItemHeap) Len() int {
	return len(h)
}

func (h ItemHeap) Less(i, j int) bool {
	return h[i].Expiration < h[j].Expiration
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *ItemHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.Index = n
	*h = append(*h, item)
}

func (h *ItemHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*h = old[0 : n-1]
	return item
}
