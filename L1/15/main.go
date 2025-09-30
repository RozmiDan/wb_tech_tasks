package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Ответы:
1) justString это срез на 100 byte на оригинальную строку и пока используется justString вся строка
не может быть очищена GC
2) Если в строке не только ASCII, то можем "отрезать" символ
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = safePrefixRunes(v, 100)
}

func safePrefixRunes(s string, n int) string {
	if n <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) > n {
		r = r[:n]
	}
	return string(r)
}

func createHugeString(size int) string {
	resStr := strings.Builder{}
	resStr.Grow(size)

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 0; i < size; i++ {
		resStr.WriteRune(letters[rand.Intn(len(letters))])
	}

	return resStr.String()
}

func main() {
	someFunc()
	fmt.Println(len(justString), justString[:10])
}
