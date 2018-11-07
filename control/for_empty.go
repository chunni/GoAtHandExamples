package main

func main() {
	sum := 0
	i := 0
	for {
		sum += i
		i++
		if i == 10 {
			break
		}
	}
	println("sum:", sum)
}
