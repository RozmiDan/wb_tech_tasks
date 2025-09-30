package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	res := append(left, pivot)
	res = append(res, right...)
	return res
}

func main() {
	arr := []int{10, 5, 2, 3, 8, 1, 7, 6, 9, 4}
	fmt.Println("before:", arr)
	sorted := quickSort(arr)
	fmt.Println("after: ", sorted)
}
