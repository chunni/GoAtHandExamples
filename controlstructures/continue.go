package main

func main() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i%2 == 0 || j%2 == 0 {
				continue
			}
			println("i", i, ",j", j)
		}
	}
}
