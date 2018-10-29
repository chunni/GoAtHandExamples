package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(n int) {
			println("start process ", n)
			time.Sleep(time.Second)
			println("end process ", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	println("All go routines done")
}
