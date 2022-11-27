package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("foo", os.O_RDONLY, 0644)
	file, err = os.OpenFile("foo", os.O_RDONLY, 0o644)
	_, _ = file, err
	
	fmt.Println(1_000_000_000)
	fmt.Println(0b00_00_01)
}
