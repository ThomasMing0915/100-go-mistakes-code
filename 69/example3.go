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

		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()

	go func() {
		defer wg.Done()

		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()

	wg.Wait()
}
