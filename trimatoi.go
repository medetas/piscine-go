package piscine

func StrLenn(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a + 1

}

func Convertt(s rune) int {
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

func Countt(s []rune) int {
	sum := 0
	m := 1
	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			sum = sum + Convertt(letter)*m
			m = m * 10
		} else {
			continue
		}
	}
	return sum
}

func TrimAtoi(s string) int {
	check := 0
	sign := 1
	var slice []rune
	if s == "" {
		return 0
	}
	for ind := range s {
		if s[ind] >= '0' && s[ind] <= '9' {
			check++
		} else if s[ind] == '-' && check == 0 {
			sign = -1
		}
	}

	slice = []rune(s)

	a := StrLenn(string(slice))

	slice2 := []rune(s)
	for index, letter := range slice {
		slice2[a-index-1] = letter
	}
	result := Countt(slice2)
	return result * sign
}
