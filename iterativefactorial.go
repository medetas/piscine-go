package main

import "fmt"

func IterativeFactorial(nb int) int {
	res := 1
	for i := 1; i <= nb; i++ {
		res = res * i
	}
	return res
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
