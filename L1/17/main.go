package main

func binSearch(arr []int, searchElem int) int {
	resIndex := -1

	var lftBound, rghtBound = 0, len(arr) - 1

	for lftBound <= rghtBound {
		midBound := (rghtBound + lftBound) / 2
		if arr[midBound] == searchElem {
			return midBound
		} else if arr[midBound] < searchElem {
			lftBound = midBound + 1
		} else {
			rghtBound = midBound - 1
		}
	}

	return resIndex
}

func main() {
	arr := []int{2, 3, 6, 8, 9, 12, 16, 21, 43, 221, 432, 543, 644, 1234}
	println(binSearch(arr, 16))
	println(binSearch(arr, 2))
	println(binSearch(arr, 1234))
	println(binSearch(arr, 1000))
}
