package main

func main() {
	var i32 int32
	var i64 int64

	i := 2 //int, inferred by the value
	i32 = 2
	i64 = 3

	//println("i == i32?", i == i32) //invalid operation: i == i32 (mismatched types int and int32)
	println("i == i32?", i == int(i32))
	println("i32 + i64 = ", i32+int32(i64))
}
