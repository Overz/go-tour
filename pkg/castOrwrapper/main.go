package castOrwrapper

import "fmt"

func Cast() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println("[Cast] int: ", i)
	fmt.Println("[Cast] float64: ", f)
	fmt.Println("[Cast] uint: ", u)
}
