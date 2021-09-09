package vars

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Declaração de variaveis
func Variaveis() {
	var a = "hello 1"
	fmt.Printf("[Var] A: %s\n", a)

	var b, c int = 1, 2
	fmt.Printf("[Var] B: %d, Var C: %d\n", b, c)

	var d int
	fmt.Printf("[Var] Int Type?: %d\n", d)

	e := "hello 2"
	fmt.Printf("[Var] Atribuition E: %s\n", e)

	var a1, a2, a3 bool
	var i1 int
	// fmt.Println(i1, a1, a2, a3, "string: ", s)
	fmt.Println("[Var] autodeclared bool: ", a1, a2, a3)
	fmt.Println("[Var] autodeclared int: ", i1)

	var r, python, java = true, false, "no!"
	fmt.Println("[Var] declared and recived bool: ", r, python)
	fmt.Print("[Var] declared and recived string: ", java)
	fmt.Println()
}
