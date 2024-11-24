package main

import "fmt"

func main() {
	fmt.Println("Learning slice")

	number := []int{1, 2, 3, 4, 5}
	fmt.Println("SLice", number)
	fmt.Println("Length", len(number))
	fmt.Println("Capacity", cap(number))

	number = append(number, 1, 2, 3, 4, 5)
	fmt.Println("SLice", number)
	fmt.Println("Length", len(number))
	fmt.Println("Capacity", cap(number))

	name := make([]int, 3, 5)
	name = append(name, 1)
	name = append(name, 1)

	fmt.Println("SLice", name)
	fmt.Println("Length", len(name))
	fmt.Println("Capacity", cap(name))

	// Now capacity will double

	name = append(name, 1)
	fmt.Println("SLice", name)
	fmt.Println("Length", len(name))
	fmt.Println("Capacity", cap(name))

	// Create empty slice
	// Method 1
	a := []int{}
	fmt.Println("SLice", a)
	fmt.Println("Length", len(a))
	fmt.Println("Capacity", cap(a))

	// Method 2
	b := make([]int, 0)
	fmt.Println("SLice", b)
	fmt.Println("Length", len(b))
	fmt.Println("Capacity", cap(b))

}
