package main

import (
	"sync"
	"time"
)

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	counter := NewCounter()

	go func() {
		counter.Increment("foo")
	}()

	go func() {
		counter.Increment("bar")
	}()

	time.Sleep(time.Second)
}