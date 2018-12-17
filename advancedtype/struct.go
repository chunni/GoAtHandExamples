package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func (s student) String() string {
	return fmt.Sprintf("student[name:%s, age:%d]", s.name, s.age)
}

func main() {
	var s1 student
	fmt.Println("use var:", s1)

	s2 := student{name: "Keven", age: 10}
	s3 := student{age: 10, name: "Keven"}
	fmt.Println("named field:", s2)
	fmt.Println("named field - different order:", s3)
	fmt.Println("values equal?", s2 == s3)

	s4 := student{name: "Kate"}
	fmt.Println("Omit field:", s4)

	s5 := new(student)
	fmt.Println("allocate with 'new':", s5)

	s6 := &student{}
	fmt.Println("values equal?", *s5 == *s6)
	fmt.Println("pointers equal?", s5 == s6)

}
