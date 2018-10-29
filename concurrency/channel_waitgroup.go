package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	sem <- 1 //initialize
	sem <- 1
	sem <- 1

	logger := log.New(os.Stderr, "", log.Lmicroseconds)

	for i := 0; i < 10; i++ {
		go func(n int) {
			//waite one semaphore by receiving data from the channel, value doesn't matter
			<-sem
			logger.Println("acquired one")
			//time.Sleep(10 * time.Microsecond)
			logger.Println("print document", n)
			time.Sleep(500 * time.Microsecond)
			//release one semaphore by send data to the channel
			sem <- 1
			logger.Println("release one")
		}(i)
	}

	time.Sleep(2 * time.Second)
}
