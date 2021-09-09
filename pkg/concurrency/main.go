package concurrency

import "fmt"

func Main() {
	async()
	fmt.Println()
	channels()
	fmt.Println()
	Select()
	fmt.Println()
	Mutex()
}
