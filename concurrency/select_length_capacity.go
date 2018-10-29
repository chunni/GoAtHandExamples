package main

func main() {
	ch := make(chan int, 5)
	println("length:", len(ch), ", capacity:", cap(ch))
	ch <- 1
	ch <- 2
	println("length:", len(ch), ", capacity:", cap(ch))
	<-ch
	println("length:", len(ch), ", capacity:", cap(ch))
}
