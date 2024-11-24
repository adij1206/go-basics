package main

import "fmt"

func main() {
	fmt.Println("Learning If-else")

	x := 5

	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is smaller than 5")
	}

	y := 10

	if y > 10 {
		fmt.Println("Y is greater than 10")
	} else if y > 5 && y <= 10 {
		fmt.Println("Y is in between 5 and 10")
	} else {
		fmt.Println("Y is smaller than 5")
	}
}
