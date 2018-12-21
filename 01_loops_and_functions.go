package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	for i < 100 {
		z -= (z*z - x) / (2 * x)
		i++
	}
	return z, i
}

func sqrt2(x float64) (float64, int) {
	z := 1.0
	prev := 0.0
	i := 0
	for prev != z {
		prev = z
		z -= (z*z - x) / (2 * x)
		i++
	}
	return z, i
}

// exercise01 run exercise 01
func exercise01() {
	x := 38892.0
	res1, it1 := sqrt(x)
	res2, it2 := sqrt2(x)
	fmt.Println(res1, it1)
	fmt.Println(res2, it2)
	fmt.Println(math.Sqrt(x))
}
