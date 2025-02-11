package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("GO", 20)
	banner("G😀", 20)

	s := "G😀"
	fmt.Println("len:", len(s))

	for i, v := range s {
		fmt.Println(i, v)
	}

	b := s[0]
	c := s[1]
	fmt.Printf("%c of type %T\n", b, b)
	fmt.Printf("%c of type %T\n", c, c)
	// strings are represented by two types
	// byte (uint8)
	// rune (int32)
	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y)
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("\ngog", isPalindrome("gog"))

	fmt.Println("\ng", isPalindrome("g"))
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	// padding := (width - len(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()

}
