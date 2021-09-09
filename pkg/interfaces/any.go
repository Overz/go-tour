package interfaces

import "fmt"

func Any() {

	// Tipo any do go
	var i interface{}

	i = 40
	describeAny(i)

	i = "hello"
	describeAny(i)
}

func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
