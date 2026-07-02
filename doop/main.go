package main

import (
	"os"
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
	args := os.Args[1:]
	var a int
	var b int
	var max int
	var min int
	var sign string
	var byted_a []byte
	var byted_b []byte
	if len(args) == 3 {
		if args[1] == "*" || args[1] == "+" || args[1] == "-" || args[1] == "/" || args[1] == "%" {
			byted_a = []byte(args[0])
			byted_b = []byte(args[2])
			for _, v := range byted_a {
				if (v <= 57 && v >= 48) || v == 45 {
					continue
				} else {
					return
				}
			}
			for _, v := range byted_b {
				if (v <= 57 && v >= 48) || v == 45 {
					continue
				} else {
					return
				}
			}
			if !inRange(args[0]) || !inRange(args[2]) {
				return
			}
			a = Atoi(args[0])
			b = Atoi(args[2])
			sign = args[1]
			max = 9223372036854775807
			min = -9223372036854775808
			if sign == "+" {
				v := a + b
				if v >= max || v <= -min {
					return
				} else {
					v := a + b
					os.Stdout.WriteString(CustomItoa(v))
					os.Stdout.WriteString("\n")
				}
			} else if sign == "*" {
				if (a >= max && (b > 1 || b < -1)) || (a <= min && (b > 1 || b < 0)) {
					return
				} else {
					v := a * b
					os.Stdout.WriteString(CustomItoa(v))
					os.Stdout.WriteString("\n")
				}
			} else if sign == "/" {
				if b != 0 {
					if a >= max || a <= -min || b >= max || b <= -min {
						return
					} else {
						v := a / b
						os.Stdout.WriteString(CustomItoa(v))
						os.Stdout.WriteString("\n")
					}
				} else {
					os.Stdout.WriteString("No division by 0")
					os.Stdout.WriteString("\n")
				}
			} else if sign == "%" {
				if b != 0 {
					if a >= max || a <= -min || b >= max || b <= min {
						return
					} else {
						v := a % b
						os.Stdout.WriteString(CustomItoa(v))
						os.Stdout.WriteString("\n")
					}
				} else {
					os.Stdout.WriteString("No modulo by 0")
					os.Stdout.WriteString("\n")
				}
			} else if sign == "-" {
				if (b < 0 && a > max+b) || (b > 0 && a < min+b) {
					return
				} else {
					v := a - b
					os.Stdout.WriteString(CustomItoa(v))
					os.Stdout.WriteString("\n")
				}
			} else {
				return
			}
		}
	}
}

func inRange(s string) bool {
	maxStr := "9223372036854775807"
	minStr := "9223372036854775808" // absolute value of min
	negative := s[0] == '-'
	digits := s
	if negative {
		digits = s[1:]
	}
	limit := maxStr
	if negative {
		limit = minStr
	}
	if len(digits) > len(limit) {
		return false
	}
	if len(digits) < len(limit) {
		return true
	}
	// same length, compare digit by digit
	return digits <= limit
}

func CustomItoa(num int) string {
	// Handle the edge case for zero directly
	if num == 0 {
		return "0"
	}

	// Track if the number is negative
	isNegative := false
	if num < 0 {
		isNegative = true
		// Avoid overflow bugs on math.MinInt by using unsigned conversion
		// if you need to support the absolute minimum boundary.
		num = -num
	}

	// A 64-bit integer has at most 19 digits + 1 sign character
	var buf [20]byte
	i := len(buf)

	// Extract digits from right to left
	for num > 0 {
		i--
		buf[i] = byte('0' + (num % 10))
		num /= 10
	}

	// Append negative sign if necessary
	if isNegative {
		i--
		buf[i] = '-'
	}

	// Convert the slice section to a string
	return string(buf[i:])
}
