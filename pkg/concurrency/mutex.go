package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Mutex() {
	teste4()
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func teste4() {
	// Cria um objeto Safe
	c := SafeCounter{v: make(map[string]int)}

	// Executa o incremento de maneir async
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// Realiza o sleep para que todo o incremento seja
	// adicionado ao objeto criado acima
	time.Sleep(time.Second)
	// Printa o valor final
	fmt.Println("[Channel] Final async increment:", c.Value("somekey"))
}
