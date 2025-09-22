package main

import "fmt"

func interOfSets(fstArr, scndArr []int) []int {
	resultArr := make([]int, 0)
	mp := make(map[int]struct{})

	for _, it := range fstArr {
		mp[it] = struct{}{}
	}

	for _, it := range scndArr {
		if _, ok := mp[it]; ok {
			resultArr = append(resultArr, it)
			delete(mp, it)
		}
	}

	return resultArr
}

func main() {
	fstArr := []int{2, 7, 3, 1, 2, 2}
	scndArr := []int{8, 5, 2, 9, 1, 8, 3}
	resArr := interOfSets(fstArr, scndArr)

	fmt.Println(resArr)
}
