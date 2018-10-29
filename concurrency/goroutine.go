package main

import "time"

func main() {
	println("start main...")

	go open()

	go func() {
		println("function literal ")
	}()

	println("wait a little while")
	time.Sleep(time.Second)
	println("end of main")
}

func open() {
	println("open file")
}
