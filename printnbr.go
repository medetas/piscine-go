package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune(45)
		n = n * -1
	}
	a := n % 10
	n = n / 10
	var c rune = '0'
	if a > 0 {
		for i := 0; i < a; i++ {
			c++
		}
		PrintNbr(n)
		z01.PrintRune(c)

	}

}
