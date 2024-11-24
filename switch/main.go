package main

import "fmt"

func main() {
	fmt.Println("Learning switch ")

	number := 3

	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Sunday")
	}

	month := "January"

	switch month {
	case "March", "April", "May", "June":
		fmt.Println("Summer")
	case "July", "August", "Sepetember", "October":
		fmt.Println("Rainy")
	case "November", "December", "January", "February":
		fmt.Println("Winter")
	}

	temperature := 25

	switch {
	case temperature < 0:
		fmt.Println("Temperature Less than zero")
	case temperature >= 0 && temperature < 25:
		fmt.Println("Temperature in between zero and twenty five")
	case temperature >= 25 && temperature < 50:
		fmt.Println("Temperature in between twenty five and fifty")
	}
}
