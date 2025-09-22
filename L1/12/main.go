package main

import (
	"fmt"
)

func uniqueStrings(arr []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, s := range arr {
		set[s] = struct{}{}
	}
	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := uniqueStrings(arr)

	fmt.Println("Множество:")
	for s := range set {
		fmt.Println(s)
	}
}
