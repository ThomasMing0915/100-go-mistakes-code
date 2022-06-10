package main

import (
	"fmt"
	"sync"
)

func receiver(messageCh <-chan int, disconnectCh <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnect, return")
					return
				}
			}
		}
	}
}

func sender(messageCh chan<- int, disconnectCh chan<- struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}
}

func main() {
	wg := new(sync.WaitGroup)
	disconnectCh := make(chan struct{})
	messageCh := make(chan int, 5)

	wg.Add(2)

	go sender(messageCh, disconnectCh, wg)

	go receiver(messageCh, disconnectCh, wg)

	wg.Wait()
}
