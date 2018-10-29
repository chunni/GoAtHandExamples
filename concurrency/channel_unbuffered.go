package main

import "time"

// Example of unbuffered channel
// which can be used to sync execution
func main() {
	println("start - async")
	go process(nil)
	println("end - async")

	time.Sleep(time.Second)
	println("---------")

	sig := make(chan int)
	println("start - sync")
	go process(sig)
	<-sig
	println("end - sync")
}

func process(channel chan int) {
	println("process... ")
	if channel != nil {
		channel <- 1
	}
}
