package main

import "fmt"

func main() {
	fmt.Println("Learning map")

	studentGrades := make(map[string]int)

	studentGrades["Aditya"] = 12
	studentGrades["Manish"] = 10
	studentGrades["Abhigyan"] = 11

	fmt.Println("Grade of Aditya", studentGrades["Aditya"])

	fmt.Println("Grade of Abhigyan", studentGrades["Abhigyan"])
	studentGrades["Abhigyan"] = 15
	fmt.Println("Grade of Abhigyan", studentGrades["Abhigyan"])

	delete(studentGrades, "Manish")
	fmt.Println("Grade of Manish", studentGrades["Manish"])

	fmt.Println("Grade of Paras", studentGrades["Paras"])
	// Above line giving the result as 0 for Paras and key is not present in the map
	// It is giving wrong impression about the element. to tackle please check below code

	parasGrade, parasGradeExists := studentGrades["Paras"]
	fmt.Println("Grade of Paras", parasGrade)
	fmt.Println("Paras Exists in map", parasGradeExists)

	for index, value := range studentGrades {
		fmt.Printf("Name of student %s and Marks of student %d\n", index, value)
	}
}
