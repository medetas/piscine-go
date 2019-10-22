package piscine

/*func StrLen(str string) int {
	a := 0
	slice := []rune(str)
	for index := range slice {
		a = index
	}
	return a + 1

}*/

func Compare(a, b string) int {
	str1 := ""
	index := -1
	if a == b {
		return 0
	} else {
		for i := 0; i < StrLen(a)-StrLen(b); i++ {
			for j := 0; j < StrLen(b); j++ {
				str1 = str1 + string(a[j+i])
			}
			if str1 == b {
				index = 1
				break
			}
			str1 = ""
		}
		return index
	}
}
