package main

import "fmt"

func main() {
	fmt.Println("LEarning Error Handling")
	//ans := divide(10, 2)

	// Here we have used the error variable
	// ans, err := divide(10, 0)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if we don't want any parameter that is getting returned from the function,
	// we can use '_' identifier
	ans, _ := divide(10, 0)

	fmt.Println("Ans", ans)
}

// func divide(a, b float64) float64 {
// 	return a / b
// }

// Throwing exception from it
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot perform action with the denominator as zero")
	}

	return a / b, nil
}
