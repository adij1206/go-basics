package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Webrequest")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while making get request", err)
		return
	}

	defer res.Body.Close()

	//fmt.Println("Response ", res)
	fmt.Printf("Datatype of response %T\n", res)

	data, err1 := ioutil.ReadAll(res.Body)

	if err1 != nil {
		fmt.Println("Error while Getting body from the response", err1)
		return
	}

	fmt.Println("Data", string(data))
}
