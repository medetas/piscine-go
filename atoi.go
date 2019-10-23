package piscine

func Atoi(s string) int {
	//a := 0
	sign := 1
	var slice []rune
	if s == "" {
		return 0
	} else if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	slice = []rune(s)
	//fmt.Println(s)
	/*for _, letter := range slice {
		fmt.Println(string(letter))
	}*/

	a := StrLen(string(slice))

	slice2 := []rune(s)
	for index, letter := range slice {
		slice2[a-index-1] = letter
	}
	result := Count(slice2)
	return result * sign
}
