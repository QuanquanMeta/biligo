package calc

import "fmt"

// this is the init of the file
func init() {
	fmt.Println("calc init function")
}

func Add(x, y int) int {
	return x + y
}
