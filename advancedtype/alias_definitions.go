package main

type aliasInt = int

//cannot define new methods on non-local type int
/*func (ai aliasInt) double() {

}*/

type defInt int

func (ai defInt) double() defInt {
	return ai * 2
}

func add2(i int) int {
	return i + 2
}

func main() {
	var ai aliasInt
	ai = 1

	var di defInt
	di = 1

	add2(ai)
	//add2(di) //cannot use di (type defInt) as type int in argument to add2
	add2(int(di))
	di.double()
}
