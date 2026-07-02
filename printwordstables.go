package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		for _, r := range v {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
