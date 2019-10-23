package main

import "os"
import "github.com/01-edu/z01"

func StrLen(str string) int {
	a := 0
	slice := []string{
		str,
	}
	for _, word := range slice {
		for _, char := range word {
			char++
			a++
		}
	}
	return a

}

func Convert(s rune) int {
	a := 0
	for true {
		if s == '0' {
			break
		}
		a++
		s--
	}
	return a
}

func Count(s []rune) int {
	sum := 0
	m := 1
	slice := []rune(s)
	for _, letter := range slice {
		if letter < '0' || letter > '9' {
			sum = 0
			break
		}
		sum = sum + Convert(letter)*m
		m = m * 10
	}
	return sum
}

func BasicAtoi2(s string) int {
	a := 0
	slice := []rune(s)
	for word := range slice {
		a = word
	}

	slice2 := []rune(s)
	for index, letter := range slice {
		slice2[a-index] = letter
	}
	result := Count(slice2)
	return result
}

func ArrStrL(arr []string) int {
	res := 0
	for index := range arr {
		res = index
	}
	return res
}

func Char(a int) rune {
	ch := ' '
	for i := 0; i < a; i++ {
		ch++
	}
	return ch
}

func main() {
	if ArrStrL(os.Args) < 2 {
		z01.PrintRune(10)
	}
	upper := false
	arr := os.Args[1:]
	if arr[0] == "--upper" {
		upper = true
		arr = arr[1:]
	}
	for _, str := range arr {
		ba := BasicAtoi2(str)
		if ba >= 1 && ba <= 26 && upper {
			ch := Char(ba + 64 - 32)
			z01.PrintRune(ch)
		} else if ba >= 1 && ba <= 26 && !upper {
			ch := Char(ba + 96 - 32)
			z01.PrintRune(ch)
		} else {
			z01.PrintRune(32)
		}
	}
	z01.PrintRune(10)
}
