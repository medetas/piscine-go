package piscine

func IsAlpha(str string) bool {
	slice := []rune(str)
	b := true

	for _, letter := range slice {
		if (letter < 48 || letter > 57) && (letter < 65 || letter > 90) && (letter < 97 || letter > 122) {
			b = false
			break
		}
	}
	return b
}
