package main

import (
	"fmt"
	"strings"
)

func uniqSymbStr(str string) bool {
	curStr := []rune(strings.ToLower(str))

	mp := make(map[rune]struct{}, len(curStr))

	for _, it := range curStr {
		if _, ok := mp[it]; ok {
			return false
		}
		mp[it] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(uniqSymbStr("asdfbcx"))
	fmt.Println(uniqSymbStr("avcxa"))
	fmt.Println(uniqSymbStr(""))
	fmt.Println(uniqSymbStr("sdffff"))
	fmt.Println(uniqSymbStr("ыавыиеку"))
}
