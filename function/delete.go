package main

import "fmt"

func main() {
	m := map[string]int{
		"Annie": 89,
		"Kyle":  91,
		"Tim":   85,
	}

	delete(m, "Kyle")
	fmt.Println("after delete:", m)

	delete(m, "Not exist")
	fmt.Println("after delete:", m)
}
