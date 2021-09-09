package matriz

import (
	"fmt"
	"strings"
)

func Matriz() {
	teste()
	teste2()
	teste3()
	// teste4()
	teste5()
}

// Declaração de um vetor e valores
// https://go-tour-br.appspot.com/moretypes/8
// matrizes subjacentes sofrem alteração de sua matriz principal
func teste() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("[Matrizes] positions: 0 -", a[0], "1 -", a[1])
	fmt.Println("[Matrizes] full string lenght:", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("[Matrizes] full integer lenght:", primes)

	// Declaração de um vetor
	// pegando do range 1 ao 4 -> [1:4]
	// algo como um Slice
	var s []int = primes[1:4]
	fmt.Println("[Matrizes] range integer [1:4]:", s)
}

func teste2() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("[Matrizes] struct headquarters:", s)
}

// Exemplo de alteração de matrizes subjacentes
func teste3() {
	fmt.Println()
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("[Matrizes] main headquarters", s)
	printSlice(s)

	s = s[1:4]
	fmt.Println("[Matrizes] sujacente 1:", s)
	printSlice(s)

	s = s[:2]
	fmt.Println("[Matrizes] sujacente 2:", s)
	printSlice(s)

	s = s[1:]
	fmt.Println("[Matrizes] sujacente 3:", s)
	printSlice(s)
}

// Printa o comprimento da matriz e sua capacidade
func printSlice(s []int) {
	fmt.Printf("[Matrizes] len: %d cap: %d %v\n", len(s), cap(s), s)
	fmt.Println()
}

// Jogo da Velha
func teste4() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[0][2] = "X"
	board[1][0] = "O"
	board[1][2] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Ignorando o indice com o operador "_"
func teste5() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("[Matrizes] Ignorando o indiece, valor: %d\n", value)
	}
}
