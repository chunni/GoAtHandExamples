package main

type intEx int

func (i intEx) double() intEx {
	return i * 2
}

//cannot define new methods on non-local type int
/*func (i int)double()  {
	i *= 2
}*/

func main() {
	i := intEx(1)
	println("double:", i.double())
}
