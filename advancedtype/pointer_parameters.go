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

func main() {
	s := "original"
	i := 1

	byVal(s, i)
	fmt.Println("main - after byVal - s:", s, ", i:", i)

	byRef(&s, &i)
	fmt.Println("main - after byRef - s:", s, ", i:", i)
}
