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

	// strings are represented by two types
	// byte (uint8)
	// rune (int32)

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
