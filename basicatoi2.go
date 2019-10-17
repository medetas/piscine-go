package main

import "fmt"

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

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := BasicAtoi2(s)
	n2 := BasicAtoi2(s2)
	n3 := BasicAtoi2(s3)
	n4 := BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
