package piscine

func StrLen(str string) int {
	a := 0
	slice := []string{
		str,
	}
	for _, word := range slice {
		for _, char := range word {
			char++
			a++
		}
	}
	return a

}
