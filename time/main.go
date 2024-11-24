package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning Time Package")

	currentTime := time.Now()
	fmt.Println("Current Time", currentTime)
	fmt.Printf("Datatype of currentTime %T\n", currentTime)

	// If we want to format the time dd-mm-yyyy all this type of syntax will not work
	// Go has specific fixed format where values of mm, yyyy, dd time and day are fixed
	// If we pass those correct value it will able to parse
	formatted := currentTime.Format("2006/02/01 Monday 15:04:05")
	fmt.Println("Formatted Time", formatted)

	// not work
	formatted = currentTime.Format("2007/02/01 Tuesday 15:04:05")
	fmt.Println("Formatted Time", formatted)

	layoutstr := "02-01-2006"
	datestr := "24-11-2024"
	time_con, _ := time.Parse(layoutstr, datestr)
	fmt.Println("COnverted string to time format", time_con)

	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("Updated Time", new_date)
}
