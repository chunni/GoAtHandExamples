package main

import "fmt"

func average(a ...int) int {
	n := len(a)

	switch n {
	case 0:
		return 0
	case 1:
		return a[0]
	default:
		sum := 0
		for _, x := range a {
			sum += x
		}
		return sum / n
	}
}

func main() {
	a := average()
	fmt.Println("average of nothing:", a)

	a = average(5)
	fmt.Println("average of a single number 5:", a)

	a = average(2, 3, 4)
	fmt.Println("average of a multiple number (2, 3, 4):", a)
}
