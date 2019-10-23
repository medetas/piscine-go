package piscine

func ToUpper(s string) string {
	slice := []rune(s)
	str := ""
	for _, letter := range slice {
		if letter >= 97 && letter <= 122 {
			letter = letter - 32
		}
		str = str + string(letter)
	}
	return str
}
