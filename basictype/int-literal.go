package main

import "fmt"

func main() {
	i := 123
	i8 := 0173
	i16 := 0x7b

	fmt.Println("123 == 0173 ?", i == i8)
	fmt.Println("0173 == 0x7b ?", i8 == i16)
	fmt.Printf("decimal: i=%d,i8=%d,i16=%d\n", i, i8, i16)
	fmt.Printf("octal: i=%o,i8=%o,i16=%o\n", i, i8, i16)
	fmt.Printf("hexadecimal: i=%x,i8=%x,i16=%x\n", i, i8, i16)
}
