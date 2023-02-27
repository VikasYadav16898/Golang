package main

import "fmt"

// Function inside function not allowed
func main() {
	fmt.Println("Welcome to functions in GoLang")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2,5,8,7)
	fmt.Println("proResult is", proResult)
	fmt.Println("proMessage is", myMessage)

}

func greeter()  {
	fmt.Println("Namaste from GoLang")
}

func proAdder(values... int) (int, string)  {
	total := 0
	for _, value := range values{
		total+= value
	}
	return total, "Hi pro result function"
}


// Value type to be passed or returned is known as function signature
func adder(valOne int, valTwo int) int  {
	return valOne + valTwo
}


// func greeterTwo()  {
// 	fmt.Println("Another Method")
// }
