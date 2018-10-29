package main

import "fmt"

const startPoint = 80

func main() {
	check(90)
	check(70)
}

func check(current int) {
	if current > startPoint {
		fmt.Println("Wow! we are going up!")
	} else {
		fmt.Println("Uh~oh~~ not good~~")
	}
}
