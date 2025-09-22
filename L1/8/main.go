package main

import (
	"fmt"
	"log"
)

func setBit(n int64, pos uint, bit bool) int64 {
	if bit {
		return n | (int64(1) << pos)
	}
	return n &^ (int64(1) << pos)
}

func main() {
	var n int64
	var pos uint
	var b int

	fmt.Print("Enter: value, bit position and bit 0|1: ")
	if _, err := fmt.Scan(&n, &pos, &b); err != nil {
		log.Fatal("scan error: ", err)
	}
	if pos > 63 || (b != 0 && b != 1) {
		log.Fatal("invalid input")
	}

	before := n
	after := setBit(n, pos, b == 1)

	fmt.Printf("before: %d (%064b)\n", before, uint64(before))
	fmt.Printf("after : %d (%064b)\n", after, uint64(after))
}
