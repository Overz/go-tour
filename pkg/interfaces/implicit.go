package interfaces

import "fmt"

type A interface {
	P()
}
type B struct {
	S string
}

// This method means type B implements the interface A,
// but we don't need to explicitly declare that it does so.
func (t B) P() {
	fmt.Println(t.S)
}

// implementação de interface implicitamente
func teste2() {
	var i A = B{"hello"}
	i.P()
}
