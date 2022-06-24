package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	// a, _ := myInteger()
	// fmt.Printf("myInteger returns as: %d\n", a)

	// myFloat()
	// myFmt()
	// myString()
	myFor()

	n, _ := chineseCount()
	fmt.Println(n)
}

// int
func myInteger() (i int, e error) {
	a := 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)
	fmt.Printf("%o \n", a)
	fmt.Printf("%x \n", a)

	i2 := 077
	fmt.Printf("%d\n", i2)

	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	fmt.Printf("%T\n", i3)

	return 0, e
}

// float
func myFloat() {
	max := math.MaxFloat32
	fmt.Printf("%f\n", max)

	f1 := 1.23456
	fmt.Printf("%f, \n", f1)

	f2 := float32(1.23456)
	fmt.Printf("%f, \n", f2)
}

// fmt
func myFmt() {
	var n = 100
	var s = "hello"
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	fmt.Printf("string: %s\n", s)
	fmt.Printf("value: %v\n", s)
	fmt.Printf("value: %#v\n", s)

}

// string utf-8 using ""
func myString() {
	path := "D:\\dev"
	fmt.Println(path)

	// raw multiple line
	s2 := `
		hello
		world	
	`
	fmt.Println(s2)

	// strin operation
	fmt.Println(len(s2))

	// stirng concatenate
	name := "li"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// contains
	fmt.Println(strings.Contains(path, "f"))

	// HasPrefix and HasSuffix
	fmt.Println(strings.HasPrefix(path, "f"))
	fmt.Println(strings.HasSuffix(path, "f"))

	s4 := "abcdef"
	fmt.Println(strings.Index(s4, "c"))

	// join
	fmt.Println(strings.Join(ret, "+"))
}

// byte and rune
func myFor() {
	s := "abcdefghijklmnopqrstuvwxyz"

	for i := range s { //byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s { //rune
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()

	// modify string. string is const
	s2 := "bailuobo" // [bailuobo]
	s3 := []rune(s2) // transfer s2 to slice of rune
	s3[0] = 'h'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	c3 := "h"
	c4 := 'h'

	fmt.Printf("c1:%T; c2:%T\n", c1, c2)
	fmt.Printf("c1:%T; c2:%T-%d\n", c3, c4, c4)
}

func chineseCount() (n int, e error) {
	s := "hello小王子"
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			n += 1
		}
	}
	return n, e
}
