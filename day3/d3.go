package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
)

var x = 100

const pi = 3.14

func init() {

	// fmt.Println(x, pi)
}

func main() {
	// myInterface()
	// myEmptyInterface()

	// pkg
	// fmt.Println("main init function")
	// ret := calc.Add(10, 20)
	// fmt.Println(ret)
	//readFromFilebyBufio()

	// readFromIoutil()

	// myOpenFileWrite()
	// myWriteFileBufio()
	// myWriteIoutil()
	// myTime()
	// myReflect()
	myStrconv()
}

// interface is a type
// interface is a type (reference type)
// interface is a type. It defines varibles have methods. It constraits types/virables that have those methods defined
// doesn't care about the type itself, but the method it can call
// A virable implements ALL the methods in the interface. Then the vriable is of that interface
// suffix with 'er' for interface
// inteface is a sets of constrain
type speaker interface {
	speak() // any variable implements speak method is a speaker
}

// n types can implement 1 interface
// 1 type can implements multiple interfaces
type cat struct{}

type dog struct{}

type person struct{}

// Use pointer receiver or value receiver
// Using pointer receiver. parameter cannot pass value
func (c *cat) speak() {
	fmt.Println("mow~")
}

func (d *dog) speak() {
	fmt.Println("wang~")
}

func (p *person) speak() {
	fmt.Println("aaa~")
}

func da(x speaker) {
	// to recieve a para of type
	x.speak() // x speaks
}

func myInterface() {
	c1 := &cat{}
	d1 := &dog{}
	p1 := &person{}

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	ss = d1 // ss has  two sections 1. type ; 2 value
	fmt.Println(ss)
}

// nil interface. All types have implemented empty interface
// interface{}
// type assertion

func show(x any) {

	switch x.(type) {
	case string:
		fmt.Printf("this is string %s\n", x)
	case int:
		fmt.Printf("this is int %d\n", x)
	default:
		fmt.Printf("%T, %v\n", x, x)
	}

}

func myEmptyInterface() {
	type any = interface{} // alias any for nil interface
	m1 := make(map[string]any, 16)
	m1["name"] = "xiang"
	m1["age"] = 9000
	m1["hobby"] = [...]string{"bascketball", "football"}
	fmt.Printf("%T, %v\n", m1, m1)

	show("abc")
	show(false)
	show(100)
	show(nil)
	show(m1)
}

// package
// package path start from GOPATH/src
// does not allow cyclic import
// alias
// anonymous import. use the init function only. not using any other methods
// func init() {}. init deos not have parameter or return
// when import the package the init function will be caled auto matically

// file
func readFromIoutil() {
	ret, err := ioutil.ReadFile("./calc/calc.go")

	if err != nil {
		fmt.Printf("read finishes, %v\n", err)
		return
	}
	fmt.Print(string(ret))
}

func readFromFilebyBufio() {
	fileObj, err := os.Open("./calc/calc.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)

	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("end of file")
			return
		}

		if err != nil {
			fmt.Printf("read finishes, %v\n", err)
			return
		}

		fmt.Print(str)
	}
}

func myFile() {
	fileObj, err := os.Open("./calc/calc.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	defer fileObj.Close()

	temp := make([]byte, 128)
	for i := 0; ; i++ {

		n, err := fileObj.Read(temp)

		if err == io.EOF {
			fmt.Println("end of file")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, err%v\n", err)
			return
		}

		fmt.Printf("%d read %d bytes\n", i, n)
		fmt.Println(string(temp[:n]))
	}

}

func myOpenFileWrite() {
	fileObj, err := os.OpenFile("./readme.md", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("open file fialed, err:%v", err)
		return
	}

	defer fileObj.Close()

	// write
	fileObj.Write([]byte("please read carefully, wirte\n"))
	// writeString
	fileObj.WriteString("WriteString method\n")

}

func myWriteFileBufio() {
	fileObj, err := os.OpenFile("./readme.md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("open file fialed, err:%v", err)
		return
	}

	defer fileObj.Close()

	w := bufio.NewWriter(fileObj)
	w.WriteString("this is an example\n")
	w.Flush() // bufio is the memory. need to flush
}

func myWriteIoutil() {
	str := "some str"
	err := ioutil.WriteFile("./readme.md", []byte(str), 0777)
	if err != nil {
		fmt.Printf("open file fialed, err:%v", err)
		return
	}
}

// logger

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter input:")
	s, _ = reader.ReadString('\n')
	fmt.Printf("your input is:%s\n", s)
}

func myLogger() {
	useBufio()
	fmt.Fprintln(os.Stdout, "this is a record of log.\n")
	fileObj, _ := os.OpenFile("./xxx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	fmt.Fprintln(fileObj, "this is a record of log.\n")
}

// time
func myTime() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())

	// timestamp
	fmt.Println(now.UnixNano())
	r := time.Unix(1232323222, 0)
	fmt.Println(r)

	fmt.Println(now.Add(24 * time.Hour))

	// timer
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// time format
	// 2022-06-26 20:30:15.000
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	// time parse
	timeObj, err := time.Parse("2006-01-02", "2022-05-21")
	if err != nil {
		fmt.Println("time parse failed")
		return
	}

	fmt.Println(timeObj)

}

// reflect
// type and kind. type cat struct. type is cat, struct is kind
//

func reflelctType(x any) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", t)
	fmt.Printf("type name:%v, type kind %v\n", t.Name(), t.Kind())
	v := reflect.ValueOf(x)

	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(200) // cannot pass in value, will panic
	}

	fmt.Printf("value kind:%v, value %v\n", v.Kind(), v)

}

func myReflect() {
	a := 3.14
	reflelctType(&a)
	b := 100
	reflelctType(&b)
	fmt.Println(b)
}

// strconv
func myStrconv() {
	str := "1000000"
	// ret1 := int64(str)

	ret1, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println("parse failed, err:", err)
		return
	}
	fmt.Printf("%v, %T\n", ret1, int(ret1))

	i1, _ := strconv.Atoi(str)
	fmt.Println(i1)
	ret3 := strconv.Itoa(i1)
	fmt.Println(ret3)

	i := int32(96)
	// ret2 := string(i) // ascii 96

	ret2 := fmt.Sprintf("%d", i)
	fmt.Println(ret2)
}
