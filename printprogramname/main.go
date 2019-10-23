package main

import "os"
import "github.com/01-edu/z01"

func main() {
	slice := []rune(os.Args[0])
	for _, letter := range slice {
		z01.PrintRune(letter)
	}
	z01.PrintRune(10)
}
