package main

import "fmt"

type s struct {
	string
	int
	//string	//duplicate field string
}

func main() {
	s := s{"a", 1}
	fmt.Println("embedded string:", s.string, ", and int:", s.int)
}
