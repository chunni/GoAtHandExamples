package main

import (
	"log"
	"os"
	"time"
)

func main() {
	sem := make(chan int, 3) //the max count of printers work at the same time
	sem <- 1                 //initialize
	sem <- 1
	sem <- 1

	logger := log.New(os.Stderr, "", log.Lmicroseconds)

	for i := 0; i < 10; i++ {
		go func(n int) {
			//waite one semaphore by receiving data from the channel, value doesn't matter
			<-sem
			logger.Println("acquired one")
			logger.Println("print document", n)
			time.Sleep(500 * time.Millisecond)
			//release one semaphore by send data to the channel
			sem <- 1
			logger.Println("release one")
		}(i)
	}

	time.Sleep(2 * time.Second)
}
