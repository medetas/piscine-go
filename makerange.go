package piscine

func MakeRange(min, max int) []int {
	var retarr []int
	if min >= max {
		return retarr
	} else {
		array := make([]int, max)
		for i := min; i < max; i++ {
			array[i] = i
		}
		retarr = array[min:]
	}
	return retarr
}
