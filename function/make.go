package main

import "fmt"

func main() {
	//s0 := make([]int) //missing len argument to make([]int)
	s := make([]int, 3)
	fmt.Println("Slice 1 - len:", len(s), ", cap:", cap(s), ", content:", s)

	s = make([]int, 1, 3)
	fmt.Println("Slice 2 - len:", len(s), ", cap:", cap(s), ", content:", s)

	m := make(map[string]int)
	m["a"] = 1
	fmt.Println("Map 1 - len:", len(m), ", content:", m)

	m = make(map[string]int, 3)
	fmt.Println("Map 2 - len:", len(m), ", content:", m)
	//m = make(map[string]int, 1, 3) //too many arguments to make(map[string]int)

	c := make(chan int)
	fmt.Println("Channel 1 - len:", len(c), ", cap:", cap(c))

	c = make(chan int, 3)
	c <- 1
	fmt.Println("Channel 2 - len:", len(c), ", cap:", cap(c))

	//c = make(chan int, 1, 3) //too many arguments to make(chan int)
}
