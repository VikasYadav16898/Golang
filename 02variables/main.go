package main

import "fmt"

const LoginToken string = "hnsfaknakjhnasknb"  //This is public

func main() {
	var userName string = "Vikas"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)

	var intVal int = 255
	fmt.Println(intVal)

	var smallFloat float32 = 255.45555555555555555555555
	fmt.Println(smallFloat)

	// Default value
	var anotherVariable int
	fmt.Println(anotherVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)

	// implicit type
	var website = "abc@gmail.com"
	fmt.Println(website)

	// No var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	// Constants
	fmt.Println((LoginToken))

}