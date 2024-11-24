package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning about Json in GO")
	person := Person{
		Name:    "Aditya",
		Age:     25,
		IsAdult: true,
	}

	fmt.Println("Person ", person)

	data, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error While Marshalling the json data", err)
		return
	}

	fmt.Println("Person data after Marshalling", string(data))

	var parsedPerson Person
	err1 := json.Unmarshal(data, &parsedPerson)

	if err1 != nil {
		fmt.Println("Error while Unmarshalling", err1)
	}

	fmt.Println("Person data after Unmarshalling", parsedPerson)
}
