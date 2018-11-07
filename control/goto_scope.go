package main

func main() {
	println("start")

	j := 10

	goto In //unresolvable label "In"

	for j < 20 {
		j++
	In:
		println("inside for")
	}

	println("other things")

}
