package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное
число горутин-воркеров.
*/

func Generator(inChan chan<- int) {
	for {
		inChan <- rand.Intn(1000)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func StartWorkers(inputCh <-chan int, countN int) {
	for i := 0; i < countN; i++ {
		go func(workerNum int) {
			for value := range inputCh {
				fmt.Printf("Worker %v received value:%v\n", workerNum, value)
			}
		}(i)
	}
}

func main() {
	fmt.Println("Starting program")

	var workersCount int
	fmt.Print("Enter workers count: ")
	if _, err := fmt.Scan(&workersCount); err != nil || workersCount <= 0 {
		fmt.Println("invalid workers count")
		return
	}

	fmt.Printf("starting program with %d workers\n", workersCount)

	in := make(chan int)

	StartWorkers(in, workersCount)

	Generator(in)
}
