package main

import "math/rand"

func main() {
	if num := rand.Intn(100); num%2 == 0 {
		println(num, "is even")
	} else {
		println(num, "is odd")
	}

	//println(num)	//compile error:  undefined: num
}
