package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s := new(Student)
	fmt.Printf("s:%v, its adress %p\n", s, s)
	fmt.Println("Name:", s.name, ", age:", s.age)
}
