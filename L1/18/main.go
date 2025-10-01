package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value atomic.Int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.value.Add(1)
}

func (c *Counter) GetValue() int {
	return int(c.value.Load())
}

func main() {
	cntr := NewCounter()
	goroutineCnt := 5
	wg := sync.WaitGroup{}

	wg.Add(goroutineCnt)
	for it := 0; it < goroutineCnt; it++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				cntr.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(cntr.GetValue())
}
