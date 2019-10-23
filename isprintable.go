package piscine

func IsPrintable(str string) bool {
	slice := []rune(str)
	b := true

	for _, letter := range slice {
		if letter < 32 || letter > 126 {
			b = false
			break
		}
	}
	return b
}
