package main

import (
	"fmt"
	"sync"
)

// Конкурентное возведение в квадрат
// Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из
// массива [2,4,6,8,10], и выведет результаты в stdout.

func Square(arr []int) <-chan int {
	wg := &sync.WaitGroup{}
	arrLen := len(arr)
	resultCh := make(chan int, arrLen)

	wg.Add(arrLen)
	for i := 0; i < arrLen; i++ {
		go func(index int) {
			defer wg.Done()
			square := arr[index] * arr[index]
			resultCh <- square
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func main() {
	arr1 := []int{2, 4, 6, 8, 10}
	arr2 := []int{2, 4, 6, 8}
	arr3 := []int{2}
	arr4 := []int{2, 4, 6, 8, 10}

	fstCh := Square(arr1)
	for val := range fstCh {
		fmt.Println(val)
	}

	scndCh := Square(arr2)
	for val := range scndCh {
		fmt.Println(val)
	}

	thdCh := Square(arr3)
	for val := range thdCh {
		fmt.Println(val)
	}

	fthCh := Square(arr4)
	for val := range fthCh {
		fmt.Println(val)
	}
}
