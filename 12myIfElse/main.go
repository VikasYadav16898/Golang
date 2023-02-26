package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in GoLang...")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular Users"
	}else if loginCount > 10{
		result = "Watchout"
	}else{
		result = "Exactly 10 login Count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less than 10")
	}

}