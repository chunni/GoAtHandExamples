package main

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "service 1 done."
	}()

	go func() {
		c2 <- "service 2 done."
	}()

	select {
	case r1 := <-c1:
		println(r1)
	case r2 := <-c2:
		println(r2)
		//if there is a default case, it will be chosen
		/*default:
		println("default")
		*/
	}

}
