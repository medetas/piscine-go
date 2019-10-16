package piscine

import "github.com/01-edu/z01"

func Recur(n int) {
	if n == 0 {
		return
	}
	c := '0'
	for i := 0; i < n%10; i++ {
		c++
	} 
	n = n / 10
	Recur(n)
	z01.PrintRune(c)
	
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune(45)
	x := -1 * n
		Recur(x)
	}else if n == 0 {
		z01.PrintRune(48)
		Recur(n)
	}else {
		Recur(n) 			
	}
}
	


