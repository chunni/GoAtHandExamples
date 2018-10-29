package main

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type worker struct {
	id int
}

func (w *worker) work(batch int) {
	println("worker", w.id, ": work batch", batch)
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return &worker{rand.Intn(50)}
		},
	}

	logger := log.New(os.Stderr, "", log.Lmicroseconds)

	for i := 0; i < 10; i++ {
		go func(n int) {
			w := pool.Get().(*worker)
			logger.Println("acquired one worker")

			w.work(n)
			time.Sleep(100 * time.Millisecond)

			pool.Put(w)
			logger.Println("put the worker back")
		}(i)
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(time.Second)
}
