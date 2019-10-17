package piscine

func Convert2(s rune) int {
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

func Count2(s []rune) int {
	sum := 0
	m := 1
	slice := []rune(s)
	for _, letter := range slice {
		if letter < '0' || letter > '9' {
			sum = 0
			break
		}
		sum = sum + Convert2(letter)*m
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
	result := Count2(slice2)
	return result
}
