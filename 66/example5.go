package main

import "fmt"

func main() {
	ch1 := make(chan int)
	close(ch1)
	// 0 0
	fmt.Println(<-ch1, <-ch1)
}
