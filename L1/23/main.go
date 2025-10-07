package main

import "fmt"

func removeElemFromSlice(arr []int, index int) []int {
	if index > len(arr)-1 || index < 0 {
		return arr
	}

	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func main() {
	fmt.Println(removeElemFromSlice([]int{12, 43, 534, 23, 3, 1}, 4))
	fmt.Println(removeElemFromSlice([]int{12, 43, 534, 23, 3, 1}, 5))
	fmt.Println(removeElemFromSlice([]int{12, 43, 534, 23, 3, 1}, 2))
	fmt.Println(removeElemFromSlice([]int{12, 43}, 4))
	fmt.Println(removeElemFromSlice([]int{12, 43, 534, 23, 3, 1}, 0))
	fmt.Println(removeElemFromSlice([]int{12}, 0))
}
