package main

import "os"
import "github.com/01-edu/z01"

func ArgLen(str []string) int {
	a := 0
	slice := str
	for index := range slice {
		a = index
	}
	return a + 1
}

func main() {
	argum := os.Args
	s := ArgLen(argum)
	for i := s - 1; i > 0; i-- {
		for _, index2 := range argum[i] {
			z01.PrintRune(index2)
		}
		z01.PrintRune(10)
	}
}
