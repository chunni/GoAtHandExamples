package main

var (
	s string
	i = 12
)

func main() {
	var f1 float32
	f1 = 12.34

	var f2 = 12.34
	f3 := 12.34 //f3 declared and not used

	var x, y int
	a, b := "", 1
	x = 30
	y = 40

	println(a, b, x, y, f1, f2)
}
