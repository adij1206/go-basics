package main

import "fmt"

func main() {
	fmt.Println("Learning pointer")

	// One way to create pointer
	var num int
	num = 4

	var ptr *int
	ptr = &num

	fmt.Println("Value of num", num)
	fmt.Println("Pointer containing address", ptr)
	fmt.Println("Pointer poiniting to value", *ptr)

	// Other way to create pointer
	a := 1
	pointer := &a

	fmt.Println("Pointer pointing to value", *pointer)

	var nilPointer *int
	if nilPointer == nil {
		fmt.Printf("Nil Pointer")
	}

	value := 5
	modifyValue(&value)
	fmt.Println("Modified value", value)
}

func modifyValue(value *int) {
	*value = *value * 2
}
