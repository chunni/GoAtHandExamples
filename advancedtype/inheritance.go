package main

import "fmt"

// struct bird is to be embedded
// it is kind of like a super class
type bird struct {
	name string
}

func (b bird) eat() {
	fmt.Println(b.name + ": I'm eating")
}

func (b bird) sing() {
	fmt.Println("sing a song")
}

//struct duck embeds bird, and has a method 'sing' of its own
type duck struct {
	bird
}

func (d duck) swim() {
	fmt.Println(d.name + ": I'm swimming")
}

func (d duck) sing() {
	fmt.Println(d.name + ": quack, quack, quack~~~")
}

//struct duck embeds bird, and has a field 'name' of its own
type lark struct {
	bird
	name string
}

func (l lark) fly() {
	fmt.Println(l.name + ": I'm flying")
}

func main() {

	d := duck{}
	d.name = "Donald"
	d.swim()
	d.sing()
	d.eat()

	l := lark{}
	l.name = "Cosette"
	l.bird.name = "Nikki"
	l.fly()
	l.eat()
}
