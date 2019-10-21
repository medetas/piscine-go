package main

import "fmt"

func IsPrime(nb int) bool {
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsPrime(0))
	fmt.Println(IsPrime(-5))
}
