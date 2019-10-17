package piscine

import "fmt"

func PrintStr(str string) {
	slice := []string{
		str,
	}
	for _, word := range slice {
		fmt.Printf(word)
	}
}
