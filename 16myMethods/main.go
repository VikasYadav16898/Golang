package main

import "fmt"

func main() {
	fmt.Println("Structs the Classes of GoLang")

	hitesh := User{"hitesh", "hitesh@go.dev", "true", 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)



}

type User struct {
	Name string
	Email string
	Status string
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}