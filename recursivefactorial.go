package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb > 1 && nb < 18 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}
