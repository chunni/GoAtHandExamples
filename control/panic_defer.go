package main

func main() {
	println("start of main")

	defer func() {
		println("defer of main")
	}()

	f1WithDefer()
	println("end of main")
}

func f1WithDefer() {
	println("start of f1")

	defer func() {
		println("defer of f1")
	}()

	f2WithDefer()
	println("end of f1")
}

func f2WithDefer() {
	println("start of f2")

	defer func() {
		println("defer of f2")
	}()

	panic("this is just a mock panic")
	println("end of f2")
}
