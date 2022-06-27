package main

import (
	"fmt"
)

func main() {
	count := 100
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main")

	// time.Sleep(time.Second)
}

// concurrent: 1 goroutine 2 channel
// goroutine. Add go in front of a function

func hello(i int) {
	fmt.Printf("hello:%d", i)
}
