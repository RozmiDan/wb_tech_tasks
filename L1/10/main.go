package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 100, -123.2, -33, -45.4}

	groups := make(map[int][]float64)
	for _, val := range temps {
		key := int(val/10) * 10
		groups[key] = append(groups[key], val)
	}

	for k, vs := range groups {
		fmt.Printf("%d: %v\n", k, vs)
	}
}
