package piscine

func PrintStr(str string) {
	slice := []string{
		str,
	}
	for _, word := range slice {
		print(word)
	}
}
