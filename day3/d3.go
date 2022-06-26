package main

import (
	"fmt"
)

func main() {
	myInterface()
}

// interface is a type
// interface is a type
// interface is a type. It defines varibles have methods. It constraits types/virables that have those methods defined
// doesn't care about the type itself, but the method it can call
// A virable implements ALL the methods in the interface. Then the vriable is of that interface
type speaker interface {
	speak() // any variable implements speak method is a speaker
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("mow~")
}

func (d dog) speak() {
	fmt.Println("wang~")
}

func (p person) speak() {
	fmt.Println("aaa~")
}

func da(x speaker) {
	// to recieve a para of type
	x.speak() // x speaks
}

func myInterface() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	ss = d1 // ss has  two sections 1. type ; 2 value
	fmt.Println(ss)
}

// package

// file
