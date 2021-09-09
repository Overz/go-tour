package ponteiros

import "fmt"

// Valores com endereço de memória
// podem ser referenciados com o operador '&'
// atribuindo literalmente o endereço para uma variavel do mesmo tipo.
//
// Também podem ser pegos os valores desses endereços
// usando o operador '*' antes da variavel.
//
// Ao alterar uma referencia em memória, tanto da variavel original
// ou da variavel que esta sendo referenciada, o valor real é alterado.
func Ponteiros() {
	// valor atribuido, e criado na memória
	value := 10

	fmt.Println("[Ponteiros] Value: ", value)
	// print do endereço em memória
	fmt.Println("[Ponteiros] Memory Address: ", &value)

	// atribuição do endereço em memoria: 0xc00001e188
	var p = &value
	fmt.Println("[Ponteiros] Recived Memory Address in var P: ", p)
	// print do valor do endereço em memória
	fmt.Println("[Ponteiros] Value of Memory Address in var P: ", *p)

	// alterando ambos os valores, pois esta alterando o endereço da memoria
	*p = 50
	fmt.Println("[Ponteiros] Changed Reference Value: ", *p)
	fmt.Println("[Ponteiros] Original Reference Value: ", value)

	fmt.Println()
	teste()
}

// Funções podem realizar um tipo de "bind" com a scruct
// que vc declara antes do nome da função, e "implementa"
// uma nova funcionalidade
type Carro struct {
	name string
}

// Não referenciado em memória
func (c Carro) andou() {
	c.name = "BMW"
	fmt.Println("[Ponteiros] 1 - Dentro da fn", c.name, "andou")
}

// Referenciado em memória
func (c *Carro) andou2() {

	c.name = "Fiesta"
	fmt.Println("[Ponteiros] 2 - Dentro da fn", c.name, "andou")
}

func teste() {
	c := Carro{
		name: "Honda",
	}

	fmt.Println("[Ponteiros] 1 - Antes da fn", c.name, "andou")
	c.andou()
	fmt.Println("[Ponteiros] 1 - Depois da fn", c.name, "andou")

	fmt.Println()

	c.name = "Fusca"
	fmt.Println("[Ponteiros] 2 - Antes da fn", c.name, "andou")
	c.andou2()
	fmt.Println("[Ponteiros] 2 - Depois da fn", c.name, "andou")
}
