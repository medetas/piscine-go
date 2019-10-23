package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argum := os.Args[1:]
	//slice := []rune(argum)
	for index := range argum {
		for _, index2 := range argum[index] {
			z01.PrintRune(index2)
		}
		z01.PrintRune(10)
	}
}
