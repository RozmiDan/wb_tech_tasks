package main

import "fmt"

func reverseStr(str string) string {
	curStr := []rune(str)

	if len(curStr) < 2 {
		return str
	}

	lftPtr, rghtPtr := 0, len(curStr)-1

	for lftPtr < rghtPtr {
		curStr[lftPtr], curStr[rghtPtr] = curStr[rghtPtr], curStr[lftPtr]
		lftPtr++
		rghtPtr--
	}

	return string(curStr)
}

func main() {
	fmt.Println(reverseStr("главрыба"))
	fmt.Println(reverseStr("главрыба") == "абырвалг")
}
