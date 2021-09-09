package deferFunc

import "fmt"

// A declaração defer adia a execução de uma função
// até o final do retorno da função.
// Os argumentos das chamadas adiadas são avaliados imediatamente,
// mas a função chamada não é executada até o retorno da função.
func Defer() {
	exemplo()
}

func exemplo() {
	fmt.Println("[Defer] Starting")

	i := 0
	for i < 3 {
		defer fmt.Println("[Defer] i: ", i)
		i++
	}

	fmt.Println("[Defer] Done")
}
