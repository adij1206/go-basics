package main

import "fmt"

/*
 * Defer is used to execute the line or function at which it has been added at the
 * last of the current method execution
 * if within the method multiple defer keyword has been used,
 * it will execute in the added function in LIFO fashion
 */

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Learning Defer")

	fmt.Println("Start of the program")

	data := sum(4, 5)
	defer fmt.Println("Sum is ", data)

	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}
