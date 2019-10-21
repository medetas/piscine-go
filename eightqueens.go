package piscine

import "github.com/01-edu/z01"

const N int = 9

var position [N]int

func Conver(n int) rune {
	charX := '0'
	for i := 0; i < n; i++ {
		charX++
	}
	return charX
}

func isSafe(qn int, rp int) bool {
	for i := 1; i < qn; i++ {
		orp := position[i]
		if orp == rp || orp == rp-(qn-i) || orp == rp+(qn-i) {
			return false
		}
	}
	return true
}

func Solve(k int) {
	if k == N {
		for i := 1; i < N; i++ {
			r := position[i]
			z01.PrintRune(Conver(r))
		}
		z01.PrintRune(10)
	} else {
		for i := 1; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				Solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	Solve(1)
}
