package piscine

import "github.com/01-edu/z01"

func ConvertInt(n int) rune {
	charX := '0'
	for i := 0; i < n; i++ {
		charX++
	}
	return charX
}

func PrintNbrInOrder(n int) {
	count := 0
	var arr []int
	if n == 0 {
		count = 1
		arr = append(arr, 0)
	}
	for n != 0 {
		arr = append(arr, n%10)
		n = n / 10
		count++
	}
	for i := count; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				min := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = min
			}
		}
	}
	for i := 0; i < count; i++ {
		z01.PrintRune(ConvertInt(arr[i]))
	}
}
