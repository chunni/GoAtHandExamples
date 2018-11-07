package main

import "fmt"

func main() {
	pi := 3.14159
	f1 := .1
	f2 := 1.
	fmt.Printf("pi*f1 = %f, pi*f2 = %f\n", pi*f1, pi*f2)

	f3 := 2E3
	f4 := 1.2345e+1
	fmt.Printf("f3=%2.2f, f4=%2.2f\n", f3, f4)

}
