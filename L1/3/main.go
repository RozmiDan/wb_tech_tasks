package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func Generator() <-chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for {
			resultCh <- rand.Intn(math.MaxInt)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	return resultCh
}

func StartWorkers(inputCh <-chan int, countN int) {
	for i := 0; i < countN; i++ {
		go func(workerNum int) {
			for {
				value := <-inputCh
				fmt.Printf("Worker %v recieve value:%v", workerNum, value)
			}
		}(i)
	}
}

func main() {
	fmt.Println("Starting programm")
	workersCount := 10
	wg := &sync.WaitGroup{}
	wg.Add(workersCount)
	StartWorkers(Generator(), workersCount)
	wg.Wait()
}
