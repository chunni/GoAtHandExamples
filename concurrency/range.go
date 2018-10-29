package main

func main() {
	data := make(chan int)
	signal := make(chan int)
	go square(data, signal)

	for i := 0; i < 10; i++ {
		data <- i
	}
	// without closing the data channel, the for loop in the square function would run forever
	close(data)

	<-signal // if line `signal <- 1` cannot be reached, there is never data to receive,
	// then, there will be a deadlock
	println("main done")
}

func square(data, signal chan int) {
	for n := range data {
		println(n * n)
	}
	println("square done")
	signal <- 1 // if data channel not being closed, this line will never been reached
}
