package main

import (
	"os"

	"github.com/01-edu/z01"
)

func insert(s string) string {
	new_s := ""
	if s[:3] == "-i=" {
		new_s += s[3:]
	} else if s[:9] == "--insert=" {
		new_s += s[9:]
	}
	return new_s
}

func main() {
	input := os.Args
	concat_s := input[len(input)-1]
	should_order := false
	new_s := ""
	insert_name := "--insert"
	insert_short := "  -i"
	insert_exp := "	 This flag inserts the string into the string passed as argument."
	order_name := "--order"
	order_short := "  -o"
	order_exp := "	 This flag will behave like a boolean, if it is called it will order the argument."
	if len(input) > 1 {
		if input[1] != "-h" && input[1] != "--help" {
			for _, inp := range input {
				if len(inp) > 3 {
					if len(inp) > 9 || inp[0:3] == "-i=" {
						insert_value := insert(inp)
						concat_s += insert_value
					}
				}
				if inp == "--order" || inp == "-o" {
					should_order = true
				} else {
					new_s = concat_s
				}
			}
			if should_order {
				byted_res := []byte(concat_s)
				for i := 0; i < len(byted_res); i++ {
					for j := 0; j < len(byted_res)-1; j++ {
						if byted_res[j] > byted_res[j+1] {
							byted_res[j], byted_res[j+1] = byted_res[j+1], byted_res[j]
						}
					}
				}
				new_s = string(byted_res)
			}
			for _, i := range new_s {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
		} else {
			for _, i := range insert_name {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			for _, i := range insert_short {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			for _, i := range insert_exp {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			for _, i := range order_name {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			for _, i := range order_short {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			for _, i := range order_exp {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
		}
	} else {
		for _, i := range insert_name {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for _, i := range insert_short {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for _, i := range insert_exp {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for _, i := range order_name {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for _, i := range order_short {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for _, i := range order_exp {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
