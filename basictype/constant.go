package main

const (
	n            = 3
	f    float32 = 3
	name         = "constant"
)

const f1 = n / 2

func main() {
	//const no			 //missing value in const declaration
	//const r = rand.Int() //rand.Int() is not a constant
	//f = 12.34            //cannot assign to f
	//const i int = "a"    //cannot convert "a" (type untyped string) to type int
	const a, b, c = 1, 2, 3
}
