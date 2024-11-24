package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning about url")

	myurl := "https://jsonplaceholder.typicode.com/todos/1?title=Aditya"

	parsedUrl, err := url.Parse(myurl)

	if err != nil {
		fmt.Println("Error while parsing the url", err)
		return
	}

	fmt.Println("Scheme", parsedUrl.Scheme)
	fmt.Println("HOst", parsedUrl.Host)
	fmt.Println("Path", parsedUrl.Path)
	fmt.Println("RAwQuery", parsedUrl.RawQuery)

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "title=Adi"

	newUrl := parsedUrl.String()
	fmt.Println("New Url", newUrl)
}
