package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	myFlags()
}

func osArgs() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args[0], os.Args[2])
	fmt.Printf("%T\n", os.Args)
}

//
func myFlags() {
	// // to create flags
	name := flag.String("name", "xiang", "pls input name:")
	age := flag.Int("age", 9000, "pls input age:")
	married := flag.Bool("married", false, "Are you married?")
	cTime := flag.Duration("ct", time.Second, "how long")

	// var name string
	// flag.StringVar(&name, "name", "xiang", "pls input name:")

	// use flag
	flag.Parse()
	fmt.Println(*name)    // get the value from address
	fmt.Println(*age)     // get the value from address
	fmt.Println(*married) // get the value from address
	fmt.Println(*cTime)   // get the value from address

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
