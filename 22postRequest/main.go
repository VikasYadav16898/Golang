package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Post Request...")

	// PerformPostJsonRequest()
	performPostFormRequest()
}

func PerformPostJsonRequest()  {
	const myUrl= "http://localhost:3000/post"

	// fake json Payload
	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with go lang",
		"price": 0,
		"platform": "learn code online.in"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostFormRequest()  {
	const myUrl= "http://localhost:3000/postform"

	// fake formdata
	data := url.Values{}
	data.Add("firstname", "vikas")
	data.Add("lastname", "yadav")
	data.Add("email", "vikas@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}