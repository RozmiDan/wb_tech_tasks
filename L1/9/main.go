package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fstStage(arr []int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)

		ticker := time.NewTicker(time.Duration(rand.Intn(1000)) * time.Millisecond)
		defer ticker.Stop()

		i := 0
		for range ticker.C {
			if i >= len(arr) {
				return
			}
			outCh <- arr[i]
			i++
		}
	}()

	return outCh
}

func scndStage(inputCh <-chan int) <-chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)

		for val := range inputCh {
			outCh <- val * 2
		}
	}()

	return outCh
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 62, 1, 2, 45, 3, 5, 3, 22, 34, 65, 334, 41, 432, 41, 23, 423, 1}

	resCh := scndStage(fstStage(arr))

	for it := range resCh {
		fmt.Println(it)
	}
}
