package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	message := [15]rune{
		'x', ' ', '=',
		' ', '0' + 42/10, '0' + 42%10,
		',', ' ', 'y',
		' ', '=', ' ',
		'0' + 21/10, '0' + 21%10, '\n',
	}

	for _, ch := range message {
		z01.PrintRune(ch)
	}
}
