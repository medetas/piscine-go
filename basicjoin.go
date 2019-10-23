package piscine

func StrLennn(str []string) int {
	a := 0
	//slice := []rune(str)
	for index := range str {
		a = index
	}
	return a + 1
}

func BasicJoin(strs []string) string {
	s := ""
	for i := 0; i < StrLennn(strs); i++ {
		s = s + strs[i]
	}
	return s
}
