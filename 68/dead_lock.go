package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *Customer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func main() {
	var c Customer

	c.UpdateAge(-1)

}
