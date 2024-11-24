package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Employee struct {
	personalDetails Person
	contactDetails  Contact
	address         Address
}

type Contact struct {
	email        string
	mobilenumber string
}

type Address struct {
	pincode string
	city    string
	state   string
}

func main() {
	fmt.Println("Learning Struct")

	var person1 Person
	person1.firstName = "Aditya"
	person1.lastName = "Jain"
	person1.age = 25

	fmt.Println("Person1", person1)

	// Other way to create struct
	person2 := Person{
		firstName: "Aditya",
		lastName:  "Jain",
		age:       24,
	}

	fmt.Println("Person 2", person2)

	// Another way to initialize the struct

	var person3 = new(Person)
	person3.firstName = "Abhigyan"
	person3.lastName = "Prakash"
	person3.age = 24

	fmt.Println(person3)

	var employee Employee

	employee.personalDetails = Person{
		firstName: "Aditya",
		lastName:  "Jain",
		age:       34,
	}

	employee.contactDetails = Contact{
		email:        "abc@gmail.com",
		mobilenumber: "8989898989",
	}

	employee.address.city = "RJN"
	employee.address.state = "CG"
	employee.address.pincode = "491441"

	fmt.Println(employee)
}
