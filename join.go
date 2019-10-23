package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i := 0; i < StrLennn(strs); i++ {
		if i != StrLennn(strs)-1 {
			s = s + strs[i] + sep
		} else {
			s = s + strs[i]
		}
	}
	return s
}
