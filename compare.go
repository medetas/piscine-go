package piscine

/*func StrLen(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a + 1

}*/

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
