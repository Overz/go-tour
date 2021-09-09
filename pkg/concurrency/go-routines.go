package concurrency

import (
	"fmt"
	"time"
)

func async() {
	fn := func(s string) {
		for i := 0; i < 2; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}

	// Ao colocar "go" na frente de uma função
	// automagicamente ela se torna "async"
	fmt.Println("[Go Routines]")
	go fn("[Go Routines] world")
	fn("[Go Routines] hello")
	fmt.Println("[Go Routines]")
}
