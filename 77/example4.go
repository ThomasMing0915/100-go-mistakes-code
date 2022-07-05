package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID int
	time.Time
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int
		Time time.Time
	}{
		ID:   e.ID,
		Time: e.Time,
	})
}

func main() {
	event := Event{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
