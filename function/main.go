package main

import "fmt"

func main() {
	fmt.Println("FUnctions")
	voidFuncWithoutParams()
	voidFunctionwithParams("Aditya")
	a := nonVoidFunctionWithoutParams()
	fmt.Println(a)
	b := nonVoidFunctionWithParams(6, "Aditya")
	fmt.Println(b)

	sum := sum(4, 5)
	fmt.Println("Sum", sum)

	test(4, 5, "Aditya")
}

// void function without params
func voidFuncWithoutParams() {
	fmt.Println("Void Function without params")
}

func voidFunctionwithParams(a string) {
	fmt.Println("Void Function with params", a)
}

func nonVoidFunctionWithoutParams() int {
	fmt.Println("NonVoid Function without params")
	return 4
}

func nonVoidFunctionWithParams(a int, b string) int {
	fmt.Println("NonVoid Function with params", a, b)
	return 4
}

// If we have all the params as same datatype, we can add the datatype at last variable only
func sum(a, b int) int {
	return a + b
}

func test(a, b int, c string) {
	fmt.Println("Test", a, b, c)
}
