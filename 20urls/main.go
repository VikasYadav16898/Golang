package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sankjfakjf"

func main() {
	fmt.Println("Welcome to handling URLS")
	fmt.Println(myUrl)

	// Parsing the url
	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val:= range qparams{
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)






}