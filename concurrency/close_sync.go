package main

func main() {
	sig := make(chan int)
	println("start - send data")
	go func() {
		println("process... ")
		sig <- 1
	}()
	<-sig
	println("end - send data")

	println("---------")

	println("start - close channel")
	go func() {
		println("process... ")
		close(sig)
	}()
	<-sig
	println("end - close channel")
}
