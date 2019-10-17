package piscine

func Convert(s rune) int {
	a := 0
	for true {
		if s == '0' {
			break
		}
		a++
		s--
	}
	return a
}

func Count(s []rune) int {
	sum := 0
	m := 1
	slice := []rune(s)
	for _, letter := range slice {
		if letter < '0' || letter > '9' {
			sum = 0
			break
		}
		sum = sum + Convert(letter)*m
		m = m * 10
	}
	return sum
}

func BasicAtoi2(s string) int {
	a := 0
	slice := []rune(s)
	for word := range slice {
		a = word
	}

	slice2 := []rune(s)
	for index, letter := range slice {
		slice2[a-index] = letter
	}
	result := Count(slice2)
	return result
}
