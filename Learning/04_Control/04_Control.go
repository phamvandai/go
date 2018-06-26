package main

import "fmt"

func main() {

	/* For loop 1 - like while in C */
	var i int = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	/* For loop 2 - No need () like C */
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/* If - No need parentheses () like C */
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "chan")
		} else { // else must be placed here, not bellow!!!
			fmt.Println(i, "le")
		}
	}

	/* Switch: no "break" like C */
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number")
	}
}
