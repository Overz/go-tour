package exercicios

import (
	"strings"

	"golang.org/x/tour/wc"
)

type Words struct {
	key   string
	value int
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, v := range strings.Fields(s) {
		result[v]++
	}

	return result
}

func Maps() {
	wc.Test(WordCount)
}
