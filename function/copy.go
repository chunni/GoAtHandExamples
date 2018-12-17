package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	n := copy(a, b)

	fmt.Println("the number of elements been copied:", n)
	fmt.Println("'a' after copy:", a)

	c := []int{6, 7, 8, 9}
	n = copy(a, c)

	fmt.Println("the number of elements been copied:", n)
	fmt.Println("'a' after copy:", a)

	d := []byte{'a', 'b'}
	s := "bcd"
	n = copy(d, s)
	fmt.Println("the number of elements been copied:", n)
	fmt.Println("'c' after copy:", d)

	//n = copy(s,c) //first argument to copy should be slice; have string
}
