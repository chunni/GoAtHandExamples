package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"China": "Beijing",
		"United Kingdom":           "London",
		"United States of America": "Washington, D.C."}
	for country := range m {
		fmt.Println("country:", country)
	}

	array := []string{"a", "b", "c"}
	for _, val := range array {
		fmt.Println("val:", val)
	}
}
