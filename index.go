package piscine

/*func StrLen(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a+1

}*/

func Index(s string, toFind string) int {
	str1 := ""
	index := -1
	for i := 0; i < StrLen(s)-StrLen(toFind); i++ {
		for j := 0; j < StrLen(toFind); j++ {
			str1 = str1 + string(s[j+i])
		}
		if str1 == toFind {
			index = i
			break
		}
		str1 = ""
	}
	return index

}
