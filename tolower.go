package piscine

func ToLower(s string) string {
	slice := []rune(s)
	str := ""
	for _, letter := range slice {
		if letter >= 65 && letter <= 90 {
			letter = letter + 32
		}
		str = str + string(letter)
	}
	return str
}
