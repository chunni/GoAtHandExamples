package main

func main() {
	println("start...")

	s := "first"
	i := 1
	defer func() {
		println("string:", s)
	}()

	defer func(n int) {
		println("i in deferred:", n)
	}(i)

	s = "second"
	i = 2

	println("i in main:", i)
	println("end")
}
