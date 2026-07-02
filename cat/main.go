package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	for i := 0; i < len(args); i++ {
		file, err := os.Open(args[i])
		if err != nil {
			e := "ERROR: "
			for _, i := range e {
				z01.PrintRune(i)
			}
			error := string(err.Error())
			for _, i := range error {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		defer file.Close()
		io.Copy(os.Stdout, file)
	}
}
