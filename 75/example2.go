package main

import "time"

func main() {
	ticker := time.NewTicker(time.Microsecond)
	// or
	ticker = time.NewTicker(1000 * time.Nanosecond)
}
