package piscine

import (
	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	} else {
		for row := 1; row <= y; row++ {
			for column := 1; column <= x; column++ {
				if column == 1 && (row == 1 || row == y) {
					z01.PrintRune('A')
				} else if column == x && (row == 1 || row == y) {
					z01.PrintRune('C')
				} else if row == 1 || row == y {
					z01.PrintRune('B')
				} else if column == 1 || column == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
