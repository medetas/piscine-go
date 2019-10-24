package main

import "fmt"

func AppendRange(min, max int) []int {
	var array []int
	if min >= max {
		return array
	} else {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
		return array
	}
	return array
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
