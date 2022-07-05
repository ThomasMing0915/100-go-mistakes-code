package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Time time.Time
}

func test() error {
	t := time.Now()
	event1 := Event{
		Time: t,
	}

	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	fmt.Println(event1 == event2)
	//fmt.Println(event1)
	//fmt.Println(event2)
	return nil
}

func main() {
	err := test()
	if err != nil {
		fmt.Println(err)
	}
}
