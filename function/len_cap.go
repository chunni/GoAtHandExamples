package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("Array - len:", len(a), ", cap:", cap(a), ", content:", a)

	s := make([]int, 1, 3)
	fmt.Println("Slice - len:", len(s), ", cap:", cap(s), ", content:", s)

	m := make(map[string]int, 3)
	m["a"] = 1
	fmt.Println("Map - len:", len(m), ", content:", m)
	//cap(m) //invalid argument m (type map[string]int) for cap

	c := make(chan int, 3)
	c <- 1
	fmt.Println("Channel - len:", len(c), ", cap:", cap(c))
}
