package main

type cart struct {
}

func (c cart) add(item string) {
	println("add an item:", item)
}

type customer struct {
	id string
}

type order struct {
	id string
	customer
	cart
}

func main() {
	c := customer{"Ruby"}
	o := order{"order 1", c, cart{}}
	o.add("item A")
	println("o.id:", o.id, ", o.customer.id:", o.customer.id)
}
