package main

import (
	"encoding/json"
	"fmt"
)

func getMessage() []byte {
	return []byte(`{"id":32,"name":"foo"}`)
}

type any interface{}

func main() {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	fmt.Printf("%T\n", m["id"])
	return
}
