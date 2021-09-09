package cleanReurn

import "fmt"

// retorna automaticamente os dois tipos
// declarados no retorno
// ao colocar "return" no final da função
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func UseSplit() {
	a, b := split(42)

	fmt.Println("[CleanReturn] A: ", a)
	fmt.Println("[CleanReturn] B: ", b)
}
