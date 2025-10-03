package main

import "fmt"

func reverseStr(curStr []rune) {
	if len(curStr) < 2 {
		return
	}

	lftPtr, rghtPtr := 0, len(curStr)-1

	for lftPtr < rghtPtr {
		curStr[lftPtr], curStr[rghtPtr] = curStr[rghtPtr], curStr[lftPtr]
		lftPtr++
		rghtPtr--
	}
}

func reverseWordsInStr(origStr string) string {
	str := []rune(origStr)

	reverseStr(str)

	var startWordInd int

	for it := 0; it <= len(str); it++ {
		if it == len(str) || str[it] == ' ' {
			reverseStr(str[startWordInd:it])
			startWordInd = it + 1
		}
	}

	return string(str)
}

func main() {
	input := "snow dog sun"
	result := reverseWordsInStr(input)
	fmt.Printf("Вход:  \"%s\"\n", input)
	fmt.Printf("Выход: \"%s\"\n", result)

	testCases := []string{
		"hello world",
		"a b c d e",
		"single",
		"",
	}

	for _, test := range testCases {
		fmt.Printf("«%s» -> «%s»\n", test, reverseWordsInStr(test))
	}
}
