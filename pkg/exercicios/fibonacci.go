package exercicios

import "fmt"

func fn() func() int {
	f2, f1 := 0, 1
	return func() int {
		f := f2
		f2, f1 = f1, f+f1

		return f
	}
}

func Fibonacci() {
	f := fn()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}
