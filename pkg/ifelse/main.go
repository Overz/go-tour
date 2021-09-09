package ifelse

import "fmt"

// NÃO EXISTE IF TERNARIO
func Ifelse() {
	if 7%2 == 0 {
		fmt.Println("[IfElse] 7 is even")
	} else {
		fmt.Println("[IfElse] 7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("[IfElse] 8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println("[IfElse] ", num, "is negative")
	} else if num < 10 {
		fmt.Println("[IfElse] ", num, "has 1 digit")
	} else {
		fmt.Println("[IfElse] ", num, "has multiple digits")
	}
}
