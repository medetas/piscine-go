package piscine

func StrRev(s string) string {
	a := 0
	slice := []rune(s)
	for word := range slice {
		a = word
	}

	slice2 := []rune(s)
	for index, letter := range slice {
		slice2[a-index] = letter
	}
	return string(slice2)
}
