// mongodb+srv://Yaduveera:<password>@cluster0.j2p1p.mongodb.net/?retryWrites=true&w=majority
// 1a2b3c4d

package main

import (
	"fmt"
	"log"
	"mongoAPI/router"
	"net/http"
)

func main()  {
	fmt.Println("Welcome to MongoDB API")
	fmt.Println("Server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000",r ))
	fmt.Println("Listening at port 4000...")
}