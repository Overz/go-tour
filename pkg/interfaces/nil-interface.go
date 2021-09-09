package interfaces

type N interface {
	M()
}

func Null() {

	// RunTime Exception
	// Não existe referencia
	// para a chamada da interface
	var n N
	describe(n)
	n.M()
}
