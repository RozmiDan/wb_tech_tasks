package main

import "fmt"

func main() {
	var fstVal int
	var scndVal int

	fmt.Println("Enter 2 numbers:")
	fmt.Scan(&fstVal, &scndVal)

	// 1-й вариант xor x = y ^ x ^ y
	// 2-й вриант fstVal, scndVal = scndVal, fstVal

	fstVal = scndVal ^ fstVal
	scndVal = scndVal ^ fstVal
	fstVal = fstVal ^ scndVal

	fmt.Println("first Value", fstVal, "second value", scndVal)
}
