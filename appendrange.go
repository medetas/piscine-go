package piscine

func AppendRange(min, max int) []int {
	var array []int
	if min >= max {
		return array
	} else if max-min < 9999999999 {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
		return array
	} else {
		return array
	}

	return array
}
