package errors

import (
	"fmt"
	"strconv"
)

func Errors() {
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("Não conseguiu converter número: %v\n", err)
	}
	fmt.Println("Inteiro convertido:", i)
}
