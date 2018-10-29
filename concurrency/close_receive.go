package main

func main() {
	b := make(chan bool)
	close(b)
	println("closed channel of bool, val:", <-b)

	i := make(chan int)
	close(i)
	println("closed channel of int, val:", <-i)

	s := make(chan string)
	close(s)
	println("closed channel of string, val:", <-s)
}
