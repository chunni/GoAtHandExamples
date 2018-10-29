package main

import "sync"

func main() {
	var singleton sync.Once
	singleton.Do(initiate)
	singleton.Do(initiate)
	singleton.Do(func() {
		println("another time")
	})
}

func initiate() {
	println("initiate")
}
