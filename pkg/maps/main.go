package maps

import "fmt"

func Maps() {
	teste()
	teste2()
}

type Vertex struct {
	Lat, Long float64
}

var a map[string]Vertex

// Criando um map
func teste() {
	a = make(map[string]Vertex)

	a["teste"] = Vertex{
		40.68433, -74.39967,
	}

	b := map[string]Vertex{
		"teste":  {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}

	fmt.Println("[Maps]", a["teste"])
	fmt.Println("[Maps]", b)
}

func teste2() {
	m := make(map[string]int)

	// Adicionando
	m["teste"] = 42
	fmt.Println("[Maps] value:", m["teste"])

	// Alterando
	m["teste"] = 48
	fmt.Println("[Maps] value:", m["teste"])

	// Excluindo
	delete(m, "teste")
	fmt.Println("[Maps] value:", m["teste"])

	// Verificando se existe com "ok"
	v, ok := m["teste"]
	fmt.Println("[Maps] value:", v, "- exists?", ok)
}
