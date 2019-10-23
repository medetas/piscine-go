package piscine

func IsLower(str string) bool {
	slice := []rune(str)
	b := true

	for _, letter := range slice {
		if letter < 97 || letter > 122 {
			b = false
			break
		}
	}
	return b
}
