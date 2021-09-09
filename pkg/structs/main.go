package structs

import "fmt"

// algo como interface
type Vertex struct {
	X int
	Y int
}

func Structs() {
	fmt.Println("[Struct] auto-declaration", Vertex{1, 2})
	teste()
}

func teste() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("[Struct] object", v.X)
}
