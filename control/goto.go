package main

func main() {
	println("start")

	i := 5
	if i < 10 {
		goto End
	}

	println("other steps")

End:
	println("end")
}
