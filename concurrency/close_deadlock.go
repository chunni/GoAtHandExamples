package main

import "time"

func main() {
	sig := make(chan int)

	go func() {
		println("process... ")
		time.Sleep(time.Second)
		sig <- 1 //panic: send on closed channel
	}()

	close(sig)
	time.Sleep(2 * time.Second)
	println("done")
}
