package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Learning File")

	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Println("Error while creating file", err)
	// 	return
	// }

	// defer file.Close()

	// content := "Hi My name is Aditya"

	// byte, err1 := io.WriteString(file, content)

	// if err1 != nil {
	// 	fmt.Println("Error while creating file", err1)
	// 	return
	// }

	// fmt.Println("Byte written", byte)

	// fmt.Println("Created SUccessfully")

	// Reading file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}

	defer file.Close()

	// reading its content
	buffer := make([]byte, 1024)

	for {
		n, err1 := file.Read(buffer)

		if err1 == io.EOF {
			break
		}

		if err1 != nil {
			fmt.Println("Error while reading the content of file", err1)
			return
		}

		fmt.Println(string(buffer[:n]))
	}

	// deprecated
	//con, err2 := ioutil.ReadFile("example.txt")
	con, err2 := os.ReadFile("example.txt")

	if err2 != nil {
		fmt.Println("Error while reading file", err2)
		return
	}

	fmt.Println(string(con))

}
