package concurrency

import "fmt"

// https://go-tour-br.appspot.com/concurrency/4
func channels() {
	teste()
	teste2()
	teste3()
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Demonstração de async chan
// que printa a soma do resultado em desordem
// simulando uma concorrencia
func teste() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("[Channel] concurrency sum:", x, y)
}

// Buffered Channel
// Serve para definir o tamanho do channel
// da mesma forma que é feito com o slice
func teste2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("[Channel] Buffered:", <-ch)
	fmt.Println("[Channel] Buffered:", <-ch)
}

func teste3() {

	// Apenas o remetente deve fechar um canal, nunca o receptor.
	// O envio em um canal fechado irá causar "panic".
	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}

		// Canais não são como arquivos, você geralmente não precisa fechá-los.
		// O encerramento só é necessário quando o receptor precisa saber que
		// não há mais valores chegando, como para terminar um laço range.
		close(c)

		v, ok := <-c
		fmt.Println("[Channel] Por que o número um, e continua true?", v, ok)
	}

	c := make(chan int, 5)
	go fibonacci(cap(c), c)

	// "for i := range c" recebe valores do canal repetidamente até que seja fechado.
	for i := range c {
		fmt.Println("[Channel] chan range and close:", i)
	}

	v, ok := <-c
	fmt.Println("[Channel] Closed Channel:", v, ok)
}
