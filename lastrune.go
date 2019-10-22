package piscine

/*func StrLen(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a + 1

}*/

func LastRune(s string) rune {
	str := []rune(s)
	if s != "" {
		return str[StrLen(s)-1]
	}
	return 0
}
