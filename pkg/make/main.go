package make

import "fmt"

// Make cria objetos com os parametros passados
func Make() {
	a := make([]int, 5)
	printSlice("[Make] a:", a)

	b := make([]int, 0, 5)
	printSlice("[Make] b:", b)

	c := b[:2]
	printSlice("[Make] c:", c)

	d := c[2:5]
	printSlice("[Make] d:", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len: %d cap: %d %v\n", s, len(x), cap(x), x)
}
