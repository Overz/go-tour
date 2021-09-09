package exercicios

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	}

	return 0, fmt.Errorf("cannor Sqrt negative number: %f", x)
}

func Errors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
