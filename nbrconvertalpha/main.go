package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	runes := []rune(s)
	LenRune := 0
	result := 0
	for i := range runes {
		LenRune = i + 1
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}

func main() {
	input := os.Args
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	ALPHABET := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(input) > 1 {
		if input[1] == "--upper" {
			for i := 2; i < len(input); i++ {
				index := Atoi(input[i])
				if index > 0 && index < 27 {
					z01.PrintRune(rune(ALPHABET[index-1]))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := 1; i < len(input); i++ {
				index := Atoi(input[i])
				if index > 0 && index < 27 {
					z01.PrintRune(rune(alphabet[index-1]))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
