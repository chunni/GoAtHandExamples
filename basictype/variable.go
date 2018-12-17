package main

var ( //in group
	s string
	i = 12
)

func main() {
	var f1 float32
	f1 = 12.34

	println("f1:", f1)

	var f2 = 12.34
	f2 = 23.45
	println("f2:", f2)
	//f3 := 12.34 //f3 declared and not used

	//multiple in one statement
	var x, y int
	x = 3
	println("x:", x, " y:", y)

	a, b := "str", 1
	println("a:", a, " b:", b)
}
