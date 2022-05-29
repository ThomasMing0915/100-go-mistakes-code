package main

import (
	"fmt"
	"sync"
	"time"
)

func G3(ch <-chan int) {
	sum := 0
	for data := range ch {
		sum += data
	}

	fmt.Println("sum", sum)
}

type rs struct {
	index int
	data  []int

	sync.Once
	sync.RWMutex
}

func (r *rs) acquireData() int {
	r.Lock()
	defer r.Unlock()

	data := r.data[r.index]
	r.index++

	return data
}

func (r *rs) isEnd() bool {
	r.Lock()
	defer r.Unlock()

	return r.index >= len(r.data)
}

func NewRs() *rs {
	rsd := &rs{
		index: 0,
		data:  make([]int, 100),
	}

	for i := 0; i < 100; i++ {
		rsd.data[i] = i + 1
	}

	return rsd
}

func G1G2(rsd *rs, ch chan<- int) {
	ps := 0

	for {
		if rsd.isEnd() {
			break
		}
		ps += rsd.acquireData()

		time.Sleep(time.Millisecond * 10)
	}

	ch <- ps

	time.Sleep(time.Second)

	rsd.Do(func() {
		close(ch)
	})
}

func main() {
	rsd := NewRs()
	ch := make(chan int, 2)

	go G1G2(rsd, ch)
	go G1G2(rsd, ch)
	go G3(ch)

	time.Sleep(time.Second * 3)
}
