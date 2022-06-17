package main

import (
	"fmt"
	"time"
)

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
					fmt.Println("case ch1 close.")
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					fmt.Println("case ch2 close.")
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()

	return ch
}

func main() {
	ch1, ch2 := make(chan int, 2), make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4

	ch := merge(ch1, ch2)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	close(ch1)
	time.Sleep(time.Second)
	close(ch2)
	time.Sleep(time.Second)

}
