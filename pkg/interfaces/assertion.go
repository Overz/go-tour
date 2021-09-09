package interfaces

import "fmt"

func Assertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// ao tentar atribuir um valo float64 para uma variavel
	// este valor não existe, lançara um erro do tipo "panic"
	f = i.(float64) // error type: "panic"
	fmt.Println(f)
}
