package piscine

func AlphaCount(str string) int {
	count := 0
	for _, letter := range str {
		if (letter <= 'Z' && letter >= 'A') || (letter <= 'z' && letter >= 'a') {
			count++
		}
	}
	return count
}
