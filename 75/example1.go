package main

import "time"

func main() {
	ticker := time.NewTicker(1000)
	for {
		select {
		case <-ticker.C:
			// do something
		}
	}
}
