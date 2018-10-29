package main

import (
	"log"
	"os"
	"time"
)

func main() {
	buffered := make(chan int, 2)
	logger := log.New(os.Stderr, "", log.Lmicroseconds)

	go func() {
		for i := 0; i < 5; i++ {
			buffered <- i
			logger.Println("send", i, "to channel")
		}
		close(buffered)
	}()

	time.Sleep(time.Second)
	for v := range buffered {
		logger.Println("receive", v, "from channel")
		time.Sleep(time.Second)
	}
}
