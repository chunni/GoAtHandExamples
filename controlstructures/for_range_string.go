package main

import (
	"fmt"
)

func main() {
	s := "Hi, 世界!"
	for idx, c := range s {
		fmt.Printf("index: %d, char: %c\n", idx, c)
	}
}
