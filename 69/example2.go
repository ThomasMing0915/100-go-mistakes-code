package main

import (
	"fmt"
	"sync"
)

func main() {
	s := make([]int, 0, 1)
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		defer wg.Done()

		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		defer wg.Done()

		s2 := append(s, 1)
		fmt.Println(s2)
	}()

	wg.Wait()
}
