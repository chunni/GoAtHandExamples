package main

import "fmt"

func main() {
	a := []int{1, 2}
	r := append(a, 3, 4)
	fmt.Println("append elements:", r)

	b := []int{5, 6}
	r = append(r, b...)
	fmt.Println("append another array:", r)
}
