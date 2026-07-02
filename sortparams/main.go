package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args
	for j := 0; j < len(input); j++ {
		for i := 1; i < len(input)-1; i++ {
			z := []byte(input[i])
			n := []byte(input[i+1])
			if z[0] > n[0] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}
	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			z01.PrintRune(rune(input[i][j]))
		}
		z01.PrintRune('\n')
	}
}
