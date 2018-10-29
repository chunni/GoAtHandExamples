package main

import (
	"log"
	"os"
	"time"
)

type printer struct {
	id     int
	logger *log.Logger
}

func (p *printer) print(args ...interface{}) {
	p.logger.Println("printer", p.id, ":", args)
	//log.Println(log.Lmicroseconds,"printer",p.id, ":", args)
}

func main() {
	//define a pool of size 3
	pool := make(chan *printer, 3)
	logger := log.New(os.Stderr, "", log.Lmicroseconds)

	//initialize
	for i := 0; i < 3; i++ {
		pool <- &printer{i, logger}
	}

	for i := 0; i < 10; i++ {
		go func(n int) {
			p := <-pool
			logger.Println("acquired printer", p.id)
			//time.Sleep(10 * time.Microsecond)
			p.print("do print job ", n)
			time.Sleep(500 * time.Millisecond)
			//release by send data to the channel
			pool <- p
			logger.Println("release printer", p.id)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
