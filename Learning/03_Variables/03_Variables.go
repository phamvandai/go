package main

import (
	"fmt"
)

// Variables declaration
// 	var x int = 1
// 	=> Can declared as: x := 1 inside function

// 	var y string = "Golang"
// 	=> Can declared as: y := "Golang" inside function

var integer_num int = 10

// Constant declaration
const str1 string = "Hello"

// Multiple variables declaration
var (
	x = 1
	y = 2
	z = "Variable"
)

func main() {
	str2 := "Golang"
	integer_num := 20

	fmt.Println(str1, str2)
	fmt.Println(integer_num)

	fmt.Println(x, y, z)
}
