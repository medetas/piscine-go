package piscine

func StrLen(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a + 1

}

func NRune(s string, n int) rune {
	str := []rune(s)
	if s != "" && n > 0 {
		return rune(str[n-1])
	} else {
		return ' '
	}
}
