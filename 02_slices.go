package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func generate(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = append(make([]uint8, dx))
		for x := range result[y] {
			result[y][x] = uint8(y - int(16*math.Sin(float64(x*2))) - 130)
		}
	}
	return result
}

// exercise02 run exercise 02
func exercise02() {
	pic.Show(generate)
}
