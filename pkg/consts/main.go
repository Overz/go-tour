package consts

import "fmt"

const s string = "Hello Constant"

func Constantes() {
	fmt.Printf("[Const] String: %s\n", s)

	const a = 1000
	fmt.Printf("[Const] A: %d\n", a)

	const b = 3e15
	fmt.Println("[Const] BigType:", b)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func ConstantesNumericas() {
	fmt.Println("[Const] 'smal' numeric const with fn", needInt(Small))
	fmt.Println("[Const] 'big' numeric const with fn", needFloat(Big))
}
