package main

import "sync"

func main() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 1
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 2
	}()
}
