package main

import "github.com/01-edu/z01"
import "os"

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	slice := []string(os.Args[1:])
	lengthOfArg := 0
	for index := range slice {
		lengthOfArg = index
	}
	if isEven(lengthOfArg) == true {
		printStr(OddMsg)
	} else {
		printStr(EvenMsg)
	}
}
