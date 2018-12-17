package main

import "fmt"

func byVal(s string, i int) {
	s = "changed"
	i = 5
	fmt.Println("byVal - s:", s, ", i:", i)
}

func byRef(s *string, i *int) {
	*s = "changed"
	*i = 5
	fmt.Println("byRef - s:", *s, ", i:", *i)
}

func changeArray(a [3]int) {
	a[0] = 4
	a[1] = 5
	fmt.Println("changeArray:", a)
}

func changeSlice(s []int) {
	s[0] = 4
	s[1] = 5
	fmt.Println("changeSlice:", s)
}

func main() {
	s := "original"
	i := 1

	byVal(s, i)
	fmt.Println("main - after byVal - s:", s, ", i:", i)

	byRef(&s, &i)
	fmt.Println("main - after byRef - s:", s, ", i:", i)

	array := [3]int{1, 2, 3}
	changeArray(array)
	fmt.Println("main - after changeArray:", array)

	slice := []int{1, 2, 3}
	changeSlice(slice)
	fmt.Println("main - after changeSlice:", slice)
}
