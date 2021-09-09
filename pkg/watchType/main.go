package watchType

import "fmt"

func WatchType() {
	a := "Hello"
	b := 10
	c := 5.432
	d := 0.876 + 0.5i

	fmt.Printf("[Type] a is of type %T\n", a)
	fmt.Printf("[Type] b is of type %T\n", b)
	fmt.Printf("[Type] c is of type %T\n", c)
	fmt.Printf("[Type] d is of type %T\n", d)
}
