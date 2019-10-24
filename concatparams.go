package piscine

func ArrStrL(arr []string) int {
	res := 0
	for index := range arr {
		res = index
	}
	return res + 1
}

func ConcatParams(args []string) string {
	str := ""
	l := ArrStrL(args)
	array := make([]string, l)
	for i := 0; i < l; i++ {
		if i == l-1 {
			array[i] = args[i]
			str = str + array[i]
		} else {
			array[i] = args[i]
			str = str + array[i] + "\n"
		}
	}
	return str
}
