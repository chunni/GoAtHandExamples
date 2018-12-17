package main

import "fmt"

func main() {
	i := 1
	s := "string"

	pi := &i
	ps := &s

	fmt.Println("pi:", pi, ", ps:", ps)

	//pi++	//invalid operation: pi++ (non-numeric type *int)

	var p *int //zero-value: nil
	fmt.Println("zero value of p:", p)

	p = pi
	fmt.Println("assign pi to p:", p, ", value of *p:", *p)

	*p = 2
	fmt.Println("i:", i)
}
