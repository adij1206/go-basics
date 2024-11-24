package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while getting data", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status received ", res.StatusCode)
		return
	}

	// method 1
	//Here we have taken the response in bytes of array, we can unmarshal the data into object
	// data, err1 := ioutil.ReadAll(res.Body)

	// if err1 != nil {
	// 	fmt.Println("Error while getting data from response", err1)
	// 	return
	// }

	// fmt.Println("Data", string(data))

	// var todo Todo
	// err2 := json.Unmarshal(data, &todo)

	// if err2 != nil {
	// 	fmt.Println("Error while Unmarshalling", err2)
	// }

	// fmt.Println("Todo Object", todo)

	var todo Todo
	err2 := json.NewDecoder(res.Body).Decode(&todo)

	if err2 != nil {
		fmt.Println("Error WHile decoding the res into json", err2)
		return
	}

	fmt.Println("Data in Todo Object", todo)
}

func performPostRequest() {
	todo := Todo{
		UserId:    123,
		Title:     "Aditya",
		Completed: true,
	}

	// Converting todo struct to Json
	jsondata, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error while parsing json")
		return
	}

	// Convert to string
	jsonstring := string(jsondata)

	reader := strings.NewReader(jsonstring)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	response, err1 := http.Post(myUrl, "application/json", reader)

	if err1 != nil {
		fmt.Println("Error while getting the response", err1)
		return
	}

	defer response.Body.Close()

	data, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		fmt.Println("Error While parsing data", err2)
		return
	}

	fmt.Println("Data", string(data))
}

func performPutRequest() {
	todo := Todo{
		UserId:    1234,
		Title:     "Aditya Jain",
		Completed: true,
	}

	// Convert to json
	jsonData, err1 := json.Marshal(todo)

	if err1 != nil {
		fmt.Println("Error WHile COnverting Todo", err1)
		return
	}

	// COnvert to String
	jsonString := string(jsonData)

	// Convert to Reader
	reader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err2 := http.NewRequest(http.MethodPut, myUrl, reader)
	if err2 != nil {
		fmt.Println("Error While crearing request", err2)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err3 := client.Do(req)

	if err3 != nil {
		fmt.Println("Error While getiing response", err3)
		return
	}

	defer res.Body.Close()

	data, err4 := ioutil.ReadAll(res.Body)

	if err4 != nil {
		fmt.Println("Error while converting response", err4)
		return
	}

	fmt.Println("Data:", string(data))
}

func performDeleteRequest() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err2 := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err2 != nil {
		fmt.Println("Error While crearing request", err2)
		return
	}

	client := http.Client{}

	res, err3 := client.Do(req)

	if err3 != nil {
		fmt.Println("Error While getiing response", err3)
		return
	}

	defer res.Body.Close()

	data, err4 := ioutil.ReadAll(res.Body)

	if err4 != nil {
		fmt.Println("Error while converting response", err4)
		return
	}

	fmt.Println("Data:", string(data))
}

func main() {
	fmt.Println("Learning CRUD OPerations...")

	performGetRequest()
	performPostRequest()
	performPutRequest()
	performDeleteRequest()
}
