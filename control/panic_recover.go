package main

func main() {
	println("start of main")

	defer func() {
		println("defer of main")
		if err := recover(); err != nil {
			println("recover of main. panic:", err.(string))
		}
	}()

	f1WithRecover()
	println("end of main")
}

func f1WithRecover() {
	println("start of f1")

	defer func() {
		println("defer of f1")
		if err := recover(); err != nil {
			println("recover of f1. panic:", err.(string))
		}
	}()

	f2WithRecover()
	println("end of f1")
}

func f2WithRecover() {
	println("start of f2")

	defer func() {
		println("defer of f2")
		if err := recover(); err != nil {
			println("recover of f2. panic:", err.(string))
		}
	}()

	panic("this is just a mock panic")
	println("end of f2")
}
