package main

import "fmt"

func main() {
	fmt.Println("Structs the Classes of GoLang")

	hitesh := User{"hitesh", "hitesh@go.dev", "true", 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)


}

type User struct {
	Name string
	Email string
	Status string
	Age int
}