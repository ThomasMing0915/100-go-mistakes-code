package main

import (
	"fmt"
	"time"
)

type Donation struct {
	balance int
	ch      chan int
}

func main() {
	donation := &Donation{ch: make(chan int)}

	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}

	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(time.Second)
			donation.balance++
			donation.ch <- donation.balance
		}
	}()

	time.Sleep(time.Second * 20)
}
