package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	return rune(str[n-1])
}
