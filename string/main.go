package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Learning Strings")

	data := "Aditya,Manish,Abhigyan,Paras,Satyam"
	parts := strings.Split(data, ",")
	fmt.Println("Splitted String", parts)

	str := "one two three four two two three"
	count := strings.Count(str, "two")
	fmt.Println("Count", count)

	str = "           Hello       World!     "
	trimStr := strings.TrimSpace(str)
	fmt.Println("trimSpace", trimStr)

	str1 := "Aditya"
	str2 := "Jain"

	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Result ", result)
}
