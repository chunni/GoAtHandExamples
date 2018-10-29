package main

import (
	"fmt"
)

func main() {
	array := []string{"a", "b", "c"}
	for idx, val := range array {
		fmt.Println("index:", idx, ", val:", val)
	}

	m := map[string]string{"China": "Beijing",
		"United Kingdom":           "London",
		"United States of America": "Washington, D.C."}
	for key, val := range m {
		fmt.Println("country:", key, ",capital:", val)
	}

	s := "Hello world!"
	for idx, c := range s {
		fmt.Printf("index: %d, char: %c\n", idx, c)
	}
}
