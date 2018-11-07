package main

import "fmt"

func main() {
	s := "abc中国字"

	length := len(s)
	fmt.Println("length:", length)

	for i := 0; i < length; i++ {
		fmt.Printf("byte %d: 0x%x, char: %c\n", i, s[i], s[i])
	}

	fmt.Println()

	ior := 0
	for i, c := range s {
		fmt.Printf("index: %d, index of rune: %d, char: %c \n", i, ior, c)
		ior++
	}
}
