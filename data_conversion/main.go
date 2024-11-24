package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Learning Dataconversion")

	var num int = 42
	fmt.Println("Value of Num is ", num)
	fmt.Printf("Datatype of Num is %T\n", num)

	//num = num +42.12

	var flotNum float64 = float64(34.4)
	fmt.Println("Value of Float Num is ", flotNum)
	fmt.Printf("Datatype of Float Num is %T\n", flotNum)

	flotNum += 13.4
	fmt.Println("Value of Float Num is ", flotNum)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Value of Str is ", str)
	fmt.Printf("Datatype of Str is %T\n", str)

	num_str := "1234"
	num_int, _ := strconv.Atoi(num_str)
	fmt.Println("Value of num_int is ", num_int)
	fmt.Printf("Datatype of num_int is %T\n", num_int)

	float_Str := "3.14"
	numer_int, _ := strconv.Atoi(float_Str)
	fmt.Println("Value of numer_int is ", numer_int)
	fmt.Printf("Datatype of numer_int is %T\n", numer_int)
	// If we will try to convert the the string value containg float type of value into int,
	//it will not be able to do it

	num_float, _ := strconv.ParseFloat(float_Str, 64)
	fmt.Println("Value of num_float is ", num_float)
	fmt.Printf("Datatype of num_float is %T\n", num_float)
}
