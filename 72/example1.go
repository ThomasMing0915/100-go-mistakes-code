package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	mu      sync.RWMutex
	balance int
}

func main() {
	donation := &Donation{}

	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}

	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()

	time.Sleep(time.Second * 20)
}
