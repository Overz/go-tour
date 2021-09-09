package exercicios

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	result := make([][]uint8, dy)
	data := make([]uint8, dx)
	for i := range result {
		for j := range data {
			data[j] = uint8((i * j) / 2)
		}
		result[i] = data
	}
	return result
}

func Slices() {
	pic.Show(Pic)
}
