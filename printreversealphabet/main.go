package main

import "github.com/01-edu/z01"

func main() {
	var a rune = 'z'
	for i := 0; i <= 25; i++ {
		z01.PrintRune(a)
		a--
	}
	z01.PrintRune(10)
}
