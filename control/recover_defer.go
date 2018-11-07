package main

func main() {
	println("start")

	if err := recover(); err != nil {
		println("recover")
	}

	f()

	println("end")
}

func f() {
	panic("panic")
}
