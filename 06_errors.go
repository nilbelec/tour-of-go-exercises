package main

import (
	"fmt"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func sqrt3(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}
	z := 1.0
	i := 0
	for i < 100 {
		z -= (z*z - x) / (2 * x)
		i++
	}
	return z, nil
}

func exercise06() {
	fmt.Println(sqrt3(2))
	fmt.Println(sqrt3(-2))
}
