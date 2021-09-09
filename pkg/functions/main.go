package functions

import (
	"fmt"
)

// https://youtu.be/vdm04bVzkLg
func Main() {
	a := fn1(5, 5)
	fmt.Println("[Functions] normal function:", a)

	b1, b2 := fn2(64)
	fmt.Println("[Functions] multiple returns", b1, b2)

	c := fn3(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("[Functions] variadic return", c)

	// Função "anonima"
	r := func(a int) {
		fmt.Println("[Functions] anonymous function", a)
	}

	r(c)

	// Função dentro de função
	r1 := func(x ...int) func() int {
		r := 0
		for v := range x {
			r += v
		}

		return func() int {
			return r
		}
	}

	fmt.Println("[Functions] funtion inside function:", r1(10, 20, 30, 40, 50))
	fmt.Println("[Functions] funtion inside function Result:", r1(10, 20, 30, 40, 50)())

	hypot := func(x, y int) int {
		return x + y
	}
	fmt.Println("[Functions] function parameter:", fn4(hypot))

	pos, neg := fn5(), fn5()
	for i := 0; i < 3; i++ {
		fmt.Println("[Functions] functions closures",
			pos(i),
			neg(-2*i),
		)
	}
}

// Normal
func fn1(a, b int) int {
	return a + b
}

// Multiplos retornos
func fn2(a int) (b, c int) {
	b = a / 10
	c = a / 20
	return
}

// Variadico
func fn3(x ...int) (r int) {
	for v := range x {
		r += v
	}
	return
}

// função valor
func fn4(fn func(a, b int) int) int {
	return fn(3, 5)
}

func fn5() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
