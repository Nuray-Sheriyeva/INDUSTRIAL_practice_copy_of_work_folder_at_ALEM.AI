package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args
	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			z01.PrintRune(rune(input[i][j]))
		}
		z01.PrintRune('\n')
	}
}
