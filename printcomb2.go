package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var x1 rune = '0'
	var y1 rune = '0'
	var z1 rune = '0'
	var w1 rune = '0'
	for x := 0; x <= 9; x++ {

		for y := 0; y <= 9; y++ {

			for z := 0; z <= 9; z++ {
				for w := 0; w <= 9; w++ {
					if z1 >= x1 && w1 >= y1 {
						z01.PrintRune(x1)
						z01.PrintRune(y1)
						z01.PrintRune(32)
						z01.PrintRune(z1)
						z01.PrintRune(w1)
						if x == 9 && y == 8 && z == 9 && w == 9 {
							continue
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
					w1++
				}
				w1 = w1 - 10
				z1++
			}
			z1 = z1 - 10
			y1++
		}
		x1++
		y1 = y1 - 10

	}
	z01.PrintRune(10)
}
