package piscine

func MakeRange(min, max int) []int {
	var retarr []int
	if min < max && max-min < 9999999999 {
		array := make([]int, max)
		for i := min; i < max; i++ {
			array[i] = i
		}
		retarr = array[min:]
	} else {
		return retarr
	}
	return retarr
}
