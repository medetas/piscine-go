package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var x1 rune = '0'
	var y1 rune = '0'
	var z1 rune = '0'
	for x := 0; x <= 9; x++ {

		for y := 0; y <= 9; y++ {

			for z := 0; z <= 9; z++ {

				if y1 > x1 && z1 > y1 {
					z01.PrintRune(x1)
					z01.PrintRune(y1)
					z01.PrintRune(z1)
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
				z1++
			}
			z1 = z1 - 10
			y1++
		}
		x1++
		y1 = y1 - 10

	}

}
