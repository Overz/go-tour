package interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

// https://go-tour-br.appspot.com/methods/10
// https://go-tour-br.appspot.com/methods/11
func Main() {
	teste()
	teste2()
}

func teste() {
	// interface
	var a Abser

	// funções normais
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat implements Abser
	a = &v // *Vertex implements Abser

	// Na linha abaixo, "v" é um Vertex de valor, e não de referencia
	// e não implementa a interface, por isso não pode ser chamado.
	// a = v

	fmt.Println(a.Abs())
}

// Implementação
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Implementação
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
