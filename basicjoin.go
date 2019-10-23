package piscine

func BasicJoin(strs []string) string {
	s := ""
	for i := 0; i < len(strs); i++ {
		s = s + strs[i]
	}
	return s
}
