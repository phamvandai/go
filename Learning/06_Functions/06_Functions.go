package main

import (
	"fmt"
)

// func func_name (variable list) return_type {
// }
func average(input []float64) float64 {
	total := 0.0
	for _, value := range input {
		total += value
	}
	return total / float64(len(input))
}

// Variadic function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func exitmsg(caller string) {
	fmt.Println(caller, "is about to quit")
	// Recovery when panic -> print panic message
	str := recover()
	fmt.Println(str)
}

func divide(x, y float64) float64 {
	if y == 0.0 {
		panic("divide by zero") // stop execution
	}
	return x / y
}

func main() {
	// Defer, auto call when exit
	defer exitmsg("main")

	data := make([]float64, 20)

	for i := 0; i < len(data); i++ {
		data[i] = float64(i / 5)
	}

	fmt.Println("Average", float64(average(data)))

	// Closure function
	closure := func(x int) int {
		return x * x
	}

	fmt.Println(closure(5))

	// Using variadic function
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))

	// generate panic
	divide(5, 0.0)

	// No execution after panic except defer
	fmt.Println("execute after panic")
}
