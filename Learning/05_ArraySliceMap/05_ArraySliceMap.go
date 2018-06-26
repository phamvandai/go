package main

import (
	"fmt"
)

func main() {

	// Array declaration
	var arr [5]int

	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	// Browse index & element value in array
	// i will be index, value is according element value
	for i, value := range arr {
		fmt.Println("arr[", i, "]=", value)
	}

	// Browse only element value in array
	var total int = 0
	for _, value := range arr {
		total = total + value
	}

	fmt.Println("Sum", total)

	// Make slice
	slice1 := make([]int, 5)

	// Declare from array
	slice2 := arr[0:5]

	// Append element with existing slice to make new
	slice3 := append(slice2, 6)

	// Copy slice
	copy(slice1, slice2)

	fmt.Println("Slice1")
	for _, value := range slice1 {
		fmt.Println(value)
	}

	fmt.Println("Slice2")
	for _, value := range slice2 {
		fmt.Println(value)
	}

	fmt.Println("Slice3")
	for _, value := range slice3 {
		fmt.Println(value)
	}

	// Map
}
