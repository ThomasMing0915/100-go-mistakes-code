package main

import "fmt"

func main() {
	ch1 := make(chan int)
	close(ch1)
	v, open := <-ch1
	// 0 false
	fmt.Println(v, open)
}
