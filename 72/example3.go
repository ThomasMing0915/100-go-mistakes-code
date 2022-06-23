package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	balance int
	ch      chan int
}

func main() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	go func() {
		for {
			time.Sleep(time.Second)
			donation.cond.L.Lock()
			donation.balance++
			donation.cond.L.Unlock()
			donation.cond.Broadcast()
		}
	}()

	time.Sleep(time.Second * 20)
}
