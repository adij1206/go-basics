package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var name string
	fmt.Println("What's your name")
	// fmt.Scan(&name)
	// fmt.Println("Your name is", name)

	// Scan method will not be able to read after it encouter first whitespace
	// For that we need to use bufio package which provide buffer read to read
	// the values from cmd

	reader := bufio.NewReader(os.Stdin)
	rename, _ := reader.ReadString('\n')
	fmt.Println("Name", rename)
}
