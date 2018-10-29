package main

func main() {
	println("start of main")
	f1()
	println("end of main")
}

func f1() {
	println("start of f1")
	f2()
	println("end of f1")
}

func f2() {
	println("start of f2")
	panic("this is just a mock panic")
	println("end of f2")
}
