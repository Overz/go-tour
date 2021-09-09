package loop

import "fmt"

func Loop() {
	i := 1
	for i <= 3 {
		fmt.Printf("[Loop] First I: %d\n", i)
		i++
	}
	fmt.Println()

	for i := 1; i < 3; i++ {
		fmt.Printf("[Loop] Second I: %d\n", i)
	}

	for {
		fmt.Println("[Loop] Infinity")
		break
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println("[Loop] Continues")
			continue
		}

		fmt.Printf("[Loop] Rest of division by two %d\n", i)
	}

	// A instrução inicial e a pós-instrução são opcionais.
	// Neste padrão se torna o while
	sum := 0
	for sum < 2 {
		sum++
		if sum == 1 {
			fmt.Println("[Loop] Opcional instructions")
		}
	}
}
