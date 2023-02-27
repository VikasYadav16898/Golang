package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video- LCO")
	PerformGetRequest()

}

func PerformGetRequest(){
	const myUrl string = "http://localhost:3000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	// Another way of converting bytedata to string format
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

	
}