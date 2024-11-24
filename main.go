package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("HI")
	myutil.Printmessage()

	var name string = "Aditya"
	fmt.Println(name)

	var number int = 10
	fmt.Println(number)

	var salary float64 = 56.8
	fmt.Println(salary)

	var randomVariable = "Aditya"
	fmt.Println(randomVariable)

	// Below is not possible in GO
	//randomVariable = 10
	//fmt.Println(randomVariable)
}
