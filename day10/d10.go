package main

import (
	"fmt"

	"github.com/biligo/day10/mysplit"
)

func main() {
	ret := mysplit.Split("ababdef", "b")
	fmt.Printf("%#v\n", ret)
}
