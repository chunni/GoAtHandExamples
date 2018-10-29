package main

func main() {
	for i := 1; i < 5; i++ {
		println("i", i)

		for j := i; j < 10; j++ {
			println("j", j)
			if j == i*2 {
				println("found. j (", j, ") is the double of i (", i, ")")
				break
			}
		}
	}
}
