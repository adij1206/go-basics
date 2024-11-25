package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Say Hello")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Say Hello Again")
}

func printHi() {
	fmt.Println("Say Hi")
}

func main() {
	fmt.Println("Learning GoRoutines")
	go printHello()
	go printHi()

	//time.Sleep(1100 * time.Millisecond)
}
