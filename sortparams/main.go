package main

import "os"
import "github.com/01-edu/z01"

func ArgLen(str []string) int {
	a := 0
	slice := str
	for index := range slice {
		a = index
	}
	return a + 1
}

func main() {
	arr := os.Args[1:]
	s := ArgLen(arr)

	for i := s; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				min := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = min
			}
		}
	}
	for index := range arr {
		for _, index2 := range arr[index] {
			z01.PrintRune(index2)
		}
		z01.PrintRune(10)
	}
}
