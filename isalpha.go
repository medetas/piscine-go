package piscine

func Capitalize(s string) string {
	slice := []rune(s)
	str := ""

	for index, letter := range slice {
		if index > 0 {
			if letter >= 97 && letter <= 122 && (slice[index-1] < 65 || slice[index-1] > 90) && (slice[index-1] < 97 || slice[index-1] > 122) && (slice[index-1] < 48 || slice[index-1] > 57) {
				letter = letter - 32
				//fmt.Println(slice[index-1])
			} else if letter >= 65 && letter <= 90 && ((slice[index-1] >= 65 && slice[index-1] <= 90) || (slice[index-1] >= 48 && slice[index-1] <= 57) || (slice[index-1] >= 97 && slice[index-1] <= 122)) {
				letter = letter + 32
			}

		} else if letter >= 97 && letter <= 122 {
			letter = letter - 32
			//fmt.Println(slice[index])
		}
		str = str + string(letter)

	}
	return str
}
