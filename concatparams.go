package piscine

func ConcatParams(args []string) string {
	str := ""
	array := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			array[i] = args[i]
			str = str + array[i]
		} else {
			array[i] = args[i]
			str = str + array[i] + "\n"
		}
	}
	return str
}
