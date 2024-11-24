package main

import "fmt"

func main() {
	fmt.Println("Learning Go")

	// Array Declaration
	var name [5]string
	name[0] = "Aditya"

	fmt.Println(name)
	// When we used %q, we let complie know that we want the values of the array with quotioned
	fmt.Printf("NAme %q \n", name)

	fmt.Println("Len", len(name))

	// Initialiaztion and declaration
	var number = [5]int{1, 2, 3, 5, 6}
	fmt.Println(number)
}
