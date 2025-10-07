package main

import (
	"fmt"
	"math/big"
)

func main() {
	// создаём большие числа
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("123456789012345678901234567890", 10)
	b.SetString("98765432109876543210", 10)

	// Результаты операций
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", mul)
	fmt.Println("a / b =", div)
}
