package main

func main() {
	println("start")

	i := 5
	if i < 10 {
		goto End //End jumps over declaration of j
	}

	j := 10

	for j < 20 {
		j++
	}

	println("other things")

End:
	println("end")
}
