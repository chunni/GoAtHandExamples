package main

import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(20 * time.Minute)
		c1 <- "service 1 done."
	}()

	go func() {
		time.Sleep(10 * time.Minute)
		c2 <- "service 2 done."
	}()

	select {
	case r1 := <-c1:
		println(r1)
	case r2 := <-c2:
		println(r2)
	case <-time.After(time.Second):
		println("time out")
	}
}
