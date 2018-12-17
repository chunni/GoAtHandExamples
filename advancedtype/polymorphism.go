package main

import (
	"fmt"
)

// iAnimal is the interface that describes an animal's daily activities
// any animal that has those activities can be kept together
type iAnimal interface {
	eat()
	play()
	sleep()
}

// struct animal is to be embedded
// it is kind of like a super class
type animal struct {
	name string
}

func (a animal) eat() {
	fmt.Println(a.name + ": I'm eating...")
}

func (a animal) play() {
	fmt.Println("let's go play...")
}

func (a animal) sleep() {
	fmt.Println(a.name + ": I'm sleeping...")
}

//struct horse embeds animal, and has methods 'play' and 'sleep' of its own
type horse struct {
	animal
}

func (h horse) play() {
	fmt.Println(h.name + ": let's play, neigh, neigh, neigh~~~")
}

func (h horse) sleep() {
	h.animal.sleep()
	fmt.Println(h.name + ": I sleep standing up.")
}

//struct cow embeds animal, and has a 'play' of its own
type cow struct {
	animal
}

func (c cow) play() {
	fmt.Println(c.name + ": let's play, moo, moo, moo~~~")
}

func keep(animals ...iAnimal) {
	for _, a := range animals {
		a.eat()
		a.play()
		a.sleep()
	}
}

func main() {
	h := horse{}
	h.name = "Gerard"

	c := cow{}
	c.name = "Nicki"

	//both h and c implement iAnimal
	keep(h, c)
}
