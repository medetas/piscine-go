package piscine

func Sqrt(nb int) int {
	s := 0
	for i := 1; i*i <= nb; i++ {
		if (nb%i == 0) && (nb/i == i) {
			s = nb / i
		} else {
			s = 0
		}
	}
	return s
}
