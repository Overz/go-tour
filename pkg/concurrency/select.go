package concurrency

import "fmt"

// Exmplo de select simples
func Select() {
	fibonacci := func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
				// Caso exista outro channels sendo usados,
				// pode haver uma "delegação" para fechar estes channels
				// chamar as funções necessárias e então retornar.
			case <-quit:
				fmt.Println("[Channel] quit")
				return

				// Uma função para transferir os dados restantes
				// pois se o case esta quit, sempre ira cair em quit
				// mas os dados não serão repassados se não chamar
				// a função para repasse.
				// neste caso:
				//
				// close(myChannel)
				// Fanout(myChannel)
				//
				// A Função QuitChannel abaixo exemplifica melhor.

			}
		}
	}

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("[Channel] async channel", <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

// uma entrada, muitas saídas
func Fannout(In <-chan int, a, b chan int) {
	for data := range In {
		select {
		case a <- data:
			fmt.Println(a)
		case b <- data:
			fmt.Println(b)
		}
	}
}

// Muitas entradas, uma saída
func Funnel(a, b <-chan int, c chan int) {
	for {
		select {
		case data, more := <-a:
			fmt.Println(data, more)
		case data, more := <-b:
			fmt.Println(data, more)

		}

		data1 := <-a
		data2 := <-b

		select {
		case c <- data1:
			fmt.Println(data1)
		case c <- data2:
			fmt.Println(data2)
		}
	}
}

// Muitas entradas e muitas saídas
func Turout(a, b <-chan int, outA, outB chan int) {
	var data int
	var more bool

	for {
		select {
		case data, more = <-a:
			fmt.Println(data, more)
		case data, more = <-b:
			fmt.Println(data, more)
		}

		select {
		// o data deste case, não é o mesmo data declarado acima
		case outA <- data:
			fmt.Println(outA)
		// o data deste case, não é o mesmo data declarado acima
		case outB <- data:
			fmt.Println(outA)
		}
	}
}

// Para evitar que os dados fiquem perdidos quando um channel for fechado
// existe o case "Quit", que chama funções para "esvaziar" os valores neste escopo
func QuitChannel(Quit <-chan int, inA, inB, outA, outB chan int) {
	var data int

	for {
		select {
		case data = <-inA:
			fmt.Println(data)
		case data = <-inB:
			fmt.Println(data)
		case <-Quit:
			close(inA)
			close(inB)

			Fannout(Quit, inA, inB)
			return
		default:
			// Para que um "block" não aconteça, um valor default deve ser
			// passado como vazio, então a função ira executar normalmente
			fmt.Println("[Channel] Default QuitChannel()")
		}
	}
}
