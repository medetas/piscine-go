package piscine

func IsUpper(str string) bool {
	slice := []rune(str)
	b := true

	for _, letter := range slice {
		if letter < 65 || letter > 90 {
			b = false
			break
		}
	}
	return b
}
