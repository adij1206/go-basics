package main

import "fmt"

func main() {
	fmt.Println("Learning loop")

	for i := 0; i < 10; i++ {
		fmt.Println("Number ", i)
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	counter := 0

	// Use of break
	for {
		fmt.Println("Infinte Loop")
		counter++

		if counter == 3 {
			break
		}
	}

	// Loop for slice
	list := []int{1, 2, 3, 4, 5}

	for index, value := range list {
		fmt.Printf("Index is %d Value is %d \n", index, value)
	}

	name := "Aditya Jain"
	for index, value := range name {
		fmt.Printf("INdex value is %d and Character value is %c \n", index, value)
	}
}
