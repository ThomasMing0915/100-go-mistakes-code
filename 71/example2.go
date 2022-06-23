package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Print("a")
	}()
	go func() {
		defer wg.Done()
		fmt.Print("b")
	}()

	wg.Wait()
}