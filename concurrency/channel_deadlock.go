package main

func main() {
	buffered := make(chan int)
	buffered <- 1
	println("put 1")
	//buffered <- 2 // fatal error: all goroutines are asleep - deadlock!

	println("pull", <-buffered) // 1
	buffered <- 3
	println("put 3")
	println("pull", <-buffered) // 3
	println("pull", <-buffered) // fatal error: all goroutines are asleep - deadlock!
}
