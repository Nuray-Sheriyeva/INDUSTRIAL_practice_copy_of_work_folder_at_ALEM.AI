package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := i; k >= '0'; k-- {
				for z := '9'; z >= '0'; z-- {
					if i == k && j <= z {
						continue
					}
					if i == '0' && j == '1' && k == '0' && z == '0' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(z)
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(z)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
