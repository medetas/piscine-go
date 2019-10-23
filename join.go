package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i := 0; i < StrLennn(strs); i++ {
		s = s + strs[i] + sep
	}
	return s
}
