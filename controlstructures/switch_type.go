package main

import "math/rand"

func main() {
	val := getValue()

	switch val := val.(type) {
	case string:
		println("string:", val)
		//fallthrough	//cannot fallthrough in type switch
	case int:
		println("int:", val)
	default:
		println("not allowed")
	}
}

func getValue() interface{} {
	i := rand.Intn(30)
	println(i)
	if i > 5 {
		return 1
	} else {
		return "s"
	}
}
