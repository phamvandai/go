package main

import (
	"fmt"
)

func main() {
	var val = 10
	var ptr *int
	ptr = &val
	fmt.Println("val", *ptr)
	fmt.Println("addr", ptr)
}
