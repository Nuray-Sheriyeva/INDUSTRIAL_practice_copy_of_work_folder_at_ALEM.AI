package main

import (
	"fmt"
	"os"
)

func stringToInt(s string) int {
	result := 0
	negative := false
	start := 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}

	if negative {
		return -result
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	count := 0
	if args[0] == "-c" {
		num := stringToInt(args[1])
		success := false
		for j := 2; j < len(args); j++ {
			file, err := os.ReadFile(args[j])
			if err != nil {
				fmt.Println(err)
				success = true
				count += 1
				continue
			}
			start := len(file) - num
			if num > len(file) {
				start = 0
			}
			if success {
				fmt.Println()
			}
			result := file[start:]
			fmt.Println("==>", args[j], "<==")
			fmt.Print(string(result))
			success = true
		}
	}
	if count != 0 {
		os.Exit(1)
	}
}
