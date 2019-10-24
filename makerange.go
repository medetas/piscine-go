package piscine

func MakeRange(min, max int) []int {
	var retarr []int
	if (min < max) && (max-min < 9999999999) {
		array := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			array[i] = i + min
		}
		retarr = array
	} else {
		return retarr
	}
	return retarr
}
