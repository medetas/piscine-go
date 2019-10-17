package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	slice := []string{
		str,
	}
	for _, word := range slice {
		for _, char := range word {
			z01.PrintRune(char)
		}
	}
}
