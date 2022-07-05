package main

import (
	"encoding/json"
	"fmt"
)

type foo struct{}

func (foo) MarshalJSON() ([]byte, error) {
	return []byte(`"foo"`), nil
}

func main() {
	b, err := json.Marshal(foo{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
