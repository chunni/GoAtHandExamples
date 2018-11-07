package main

func main() {
	println("start...")

	defer func() {
		println("deffer")
	}()

	panic("there is something wrong")

	println("end")
}
