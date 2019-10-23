package piscine

func IsNumeric(str string) bool {
	slice := []rune(str)
	b := true

	for _, letter := range slice {
		if letter < 48 || letter > 57 {
			b = false
			break
		}
	}
	return b
}
