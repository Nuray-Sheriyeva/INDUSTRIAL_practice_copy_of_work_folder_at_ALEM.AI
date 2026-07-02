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

func Itoa(n int) string {
	// Edge case: if the number is exactly 0
	if n == 0 {
		return "0"
	}

	// Track if the number is negative
	isNegative := false
	if n < 0 {
		isNegative = true
		// Note: In a production environment with int min values (-2147483648),
		// absolute inversion can overflow. For simplicity, we assume safe ranges
		// or use uint64 conversion under the hood.
		n = -n
	}

	// Slice to collect bytes in reverse order
	var bytes []byte

	// Extract digits one by one from right to left
	for n > 0 {
		digit := n % 10
		// Convert digit to its ASCII character representation ('0' is 48)
		bytes = append(bytes, byte('0'+digit))
		n /= 10
	}

	// If it was a negative number, append the minus sign
	if isNegative {
		bytes = append(bytes, '-')
	}

	// Reverse the byte slice to get the correct order
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}

func main() {
	input := os.Args
	s := ""
	for i := 1; i < len(input); i++ {
		s += input[i]
		if i != len(input)-1 {
			s += " "
		}
	}
	vowels := "aeuoiAEUOI"
	var indexes []int
	lettersStr := ""
	for i, l := range s {
		for _, j := range vowels {
			if l == j {
				indexes = append(indexes, i)
				lettersStr += string(l)
			}
		}
	}
	if len(indexes) > 0 {
		letters := []rune(lettersStr)

		for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
			letters[i], letters[j] = letters[j], letters[i]
		}
		// fmt.Println(indexes)
		// fmt.Println(string(letters))
		new_s := s
		for i := 0; i < len(s); i++ {
			// fmt.Println("i: ", i, " | ")
			for j, char := range indexes {
				// fmt.Println("char: ", char, " | ")
				if i == char {
					// fmt.Println(i, " == ", char)
					new_s = new_s[:i] + string(letters[j]) + new_s[i+1:]
					// fmt.Println("SWAPPED    ", string(s[i]), "  ->    ", string(letters[j]), "   ", new_s)
				}
				// fmt.Println("________________________")
			}
		}

		for _, char := range new_s {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	} else if s == "" {
		z01.PrintRune('\n')
	} else {
		for _, char := range s {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
