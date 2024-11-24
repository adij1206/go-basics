package myutil

import "fmt"

// Public method as the first Letter is capital
func Printmessage() {
	fmt.Print("Hello world")
	printmessage()
}

// Private Method as first Letter is small
func printmessage() {
	fmt.Println("Test")
}
