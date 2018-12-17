package main

import "fmt"

type printer interface {
	print(a interface{})
}

type inkPrinter struct {
}

func (i inkPrinter) print(a interface{}) {
	fmt.Println("ink printer print:", a)
}

type laserPrinter struct {
}

func (i laserPrinter) print(a interface{}) {
	fmt.Println("laser printer print:", a)
}

type pseudo struct {
}

func hello(p printer) {
	p.print("hello world.")
}

func main() {
	ip := inkPrinter{}
	ip.print("abc")

	lp := laserPrinter{}
	lp.print(123)

	hello(ip)
	hello(lp)
	//hello(pseudo{}) //pseudo does not implement printer (missing print method)
}
