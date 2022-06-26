package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// a, _ := myInteger()
	// fmt.Printf("myInteger returns as: %d\n", a)

	// myFloat()
	// myFmt()
	// myString()
	// myFor()

	// n, _ := chineseCount()
	// fmt.Println(n)

	//myArray()

	//mySlice()

	// myPointer()
	// myMap()

	// fmt.Println(myDefer())
	// myFuncPassing()

	// myPanicRecover()

	// myStruct()
	// myConstructor()
	// myMethod()
	myJson()
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

	i3 := "0x123456"
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
func myrune() {
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

// if

func myIf() {
	if age := 19; age > 18 {
		fmt.Println("adult")
	} else if age > 35 {
		fmt.Println("midage")
	} else {
		fmt.Println("go back home do homework")
	}
}

func myFor() {
	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	// for range
	s := "hello"

	for i, v := range s {
		fmt.Printf("%d, %c\n", i, v)
	}
}

// switch
func mySwitch() {
	// switch
	switch n := 5; n {
	case 1, 2:
		fmt.Println(1)
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println(n)
	}
}

func myOperator() {

	var (
		a = 1
		b = 2
	)

	fmt.Println(a + b)

	// unary operator
	// &
	fmt.Println(5 & 2)
	// |
	fmt.Println(5 | 2)
	// ^ if different to be
	fmt.Println(5 ^ 2)

	fmt.Println(1 << 10)
}

func myArray() {
	var a1 [3]bool
	fmt.Printf("%T", a1)
	fmt.Println(a1)

	// init
	a10 := [...]int{0, 1}
	fmt.Println(a10)

	// index init
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	for i, v := range a3 {
		fmt.Println(i, v)
	}

	// multiple array
	a11 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	for i, v1 := range a11 {
		fmt.Println(i, v1)
		for j, v2 := range v1 {
			fmt.Println(j, v2)
		}
	}
}

// slice is a pointer to a block of consective memory
// cannot compare slices. the only available comparison is slice == nil
// use len(s) == 0 to determine if a slice is empty. not slice == nil
func mySlice() {
	s1 := []int{1, 2, 3, 4, 5} // define a int type slice
	s2 := []string{"hao", "han", "ge"}
	fmt.Println(s1, s2)
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d\n", len(s2), cap(s2))

	s3 := s1[0:3] // left close right open
	s4 := s1[:2]  // ->[0:2]
	s5 := s1[2:]  // -> [2:len(s1)]
	s6 := s1[:]   // -> [0, len(s1)]
	fmt.Print(s3, s4, s5, s6)
	// the cap of the slice is from the first element to the slice to the end of underlining array
	fmt.Printf("len(s4)=%d, cap(s4)=%d", len(s4), cap(s4))

	// slice is a quote
	s1[4] = 500 // modify the underline array/slice will change the slice
	fmt.Print(s3, s4, s5, s6)

	fmt.Println()
	// make
	m1 := make([]int, 5, 10)
	fmt.Println(m1)
	fmt.Printf("s1=%d, len(s4)=%d, cap(s4)=%d", m1, len(m1), cap(m1))

	// append, must use a variable to receive the return value
	a1 := []string{"aa", "bb", "cc"}
	a1 = append(a1, "dd") // The return value must the same as the first parameter, a1 since the underlining array has changed
	fmt.Printf("a1=%v len(a1)=%d, cap(a1)=%d", a1, len(a1), cap(a1))

	a2 := []string{"ff", "gg"}
	a1 = append(a1, a2...) // unpack slice
	fmt.Println(a1)

	// copy. it's a deep copy
	var a3 = make([]string, len(a1))
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = "100"
	fmt.Println(a1, a2, a3)

	// remove between 1 and 3
	// slice points to a underline array. slice does not store values
	a1 = append(a1[:1], a1[2:]...)
	fmt.Printf("a1=%v, cap(a1)%d", a1, cap(a1))
	//sort.Strings(a1) // sort
	sort.Sort(sort.Reverse((sort.StringSlice(a1)))) // reverse sort
	fmt.Println(a1)
}

// does not have pointer operations
// `&` address of
// `*` get value from address
func myPointer() {
	n := 19

	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	// new function
	var a = new(int) // allocate a new memory and return the addres to a of that type (*string, *int)
	*a = 100
	fmt.Println(*a)

	// `make` is to allocate memory for slice, map and chan and return the type itself
	b := make([]string, 5)
	fmt.Println(b)
}

// map is reference type of key/value hash table.
func myMap() {

	var m1 = make(map[string]int, 3)

	m1["good"] = 1
	m1["bad"] = 0

	v, ok := m1["empty"]
	if !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// delete to delete key/value from map
	delete(m1, "bad")

	// slice ofmap type
	var s1 = make([]map[int]string, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "aa"
	// fmt.Println(s1)
	fmt.Printf("%v", s1)

	// a map of slice
	var m2 = make(map[string][]int, 10)

	var b = make([]int, 3)
	b = []int{0, 0, 1}

	m2["aaa"] = b
	fmt.Printf("%v", m2)

}

// function. go does not have defualt parameters
func sum(x int, y int) (ret int) {
	ret = x + y
	return ret
}

// varidic function
// ... must be at last of parameters
func myVaridic(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y is a slice of []type
}

// defer (good to free file handle, database connection, socket connection)
func myDefer() (x int) {
	fmt.Println("a")
	defer fmt.Println("e") // defer the func and push it into a stack
	defer fmt.Println("d")
	defer fmt.Println("c")
	fmt.Println("b")

	x = 5
	defer func(x int) {
		x++
	}(x)

	return x // defer is in between of ret being assigned and func return
}

// func type passing
func f1(x func() int) func(int, int) int {

	// anounymous function
	var ff1 = func(x, y int) int {
		return x + y
	}

	// lambda
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	return ff1
}

func myFuncPassing() {
	a := f1
	fmt.Printf("%T\n", a)
}

// closure is a function contains a parameter from outter scope
func f3(f func()) {
	fmt.Println("This is f3")
	f()
}

func f4(x, y int) {
	fmt.Println("this is f4")
}

func f5(f func(int, int), x, y int) func() {
	tempfn := func() {
		f(x, y)
	}
	return tempfn
}

func myClosure() {
	ret := f5(f4, 100, 200)
	ret()
}

// panic & recover. recover must use together with panic
// defer need to be before panic
func pA() {
	fmt.Println("this is pA")
}

func pB() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("wrong!!!")
	fmt.Println("this is pB")
}

func pC() {
	fmt.Println("this is pC")
}

func myPanicRecover() {
	pA()
	pB()
	pC()
}

// fmt
func myFmt1() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
}

// user defined type
// type alias

type myInt int
type yourInt = int

func myType() {
	var n myInt = 100
	fmt.Printf("%T\n", n)

	var m yourInt = 100
	fmt.Printf("%T\n", m)

	var c rune = 'a'
	fmt.Printf("%T\n", c)

	var d byte = 'a'
	fmt.Printf("%T\n", d)
}

// struct is avlue type
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// struct is value type
// struct use consective memory
func myStruct() {
	// declar a person
	var p person
	p.name = "abc"
	p.age = 9000
	p.gender = "male"
	p.hobby = []string{"basketball", "football"}
	fmt.Println(p)

	// anonimous struct
	var s struct {
		x string
		y int
	}
	s.x = "aaa"
	s.y = 2
	fmt.Printf("%T, %v", s, s)

	fs1 := func(x person) {
		x.gender = "female"
	}

	fs2 := func(x *person) {
		x.gender = "female"
	}

	// pass by value VS pass by reference
	var p2 person
	p2.gender = "male"

	fs1(p2)
	fmt.Println(p2.gender)
	fs2(&p2)
	fmt.Println(p2.gender)

	// 1 init key-value
	var p3 = person{
		name:   "yuanshuai",
		gender: "male",
	}
	fmt.Println(p3)

	// 2 using the value to init
	p4 := person{
		"xiaowangzi",
		8000,
		"female",
		[]string{"basketball"},
	}
	fmt.Println(p4)

}

// 3 constructor: using prefix of new for return pointer
// using return pointer when the struct is large

type personOne struct {
	name string
	age  int
}

func newPerson(name string, age int) *personOne {
	return &personOne{
		name: name,
		age:  age,
	}
}

func myConstructor() {
	p1 := newPerson("yuanshui", 19)
	p2 := newPerson("xiang", 30)
	fmt.Println(p1, p2)
}

// Method and receiver
// Method is a func applying to a special type
type dog struct {
	Name string `json:"name",db:"name",ini:"name"`
	Age  int    `json:"age",db:"age",ini:"age"`
}

// dog ctor
func newDog(name string, age int) dog {
	return dog{
		Name: name,
		Age:  age,
	}
}

// receiver in the () before the func name
// receiver is the type. the instance is the first char of that type
func (d dog) wang() {
	fmt.Printf("%s :wangwangwang~\n", d.Name)
}

// 1. need to modify the receiver
// 2. the receiver is big
// 3. be consistent with others
func (d *dog) guonian() {
	d.Age++
}

// only applies to the user defined type and within the same pkg
func myMethod() {
	d := newDog("xiaohei", 2)
	d.wang()

	fmt.Println(d.Age)
	d.guonian()
	fmt.Println(d.Age)
}

// identifier variable , func name
// identifier the first char is captial letter. It's visible to outside pkg

// Cat this a cat struct
type Cat struct {
	name string
	dog
}

// struct - anonymous field // struct nested struct // anonymous nested struct
type myAnonymous struct {
	string
}

// struct to implement inherit
func myInherit() {

	c1 := Cat{
		name: "a cat",
		dog:  dog{"dog", 2},
	}

	c1.wang()
}

// struct and Json

func myJson() {
	d1 := dog{
		Name: "dd",
		Age:  2,
	}
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("marshal fialed, err: %v\n", err)
	}
	fmt.Printf("%v\n", string(b))

	str := `{"name":"xiang","age":18}`
	var d2 dog
	json.Unmarshal([]byte(str), &d2) // pass reference
	fmt.Printf("%#v\n", d2)
}
