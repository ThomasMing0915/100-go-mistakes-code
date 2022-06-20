package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 42
	fmt.Println(s1)
}
