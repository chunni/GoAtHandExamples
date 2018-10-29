package main

func main() {
	println("start...")

	defer func() {
		println("deffer")
	}()

	println("execute, execute, execute.")

	println("end")
}
