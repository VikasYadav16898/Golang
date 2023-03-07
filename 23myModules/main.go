package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in GoLang...")
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")


	// server
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter()  {
	fmt.Println("Hey there mod user")	
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to GoLang series on Youtube.</h1>"))
}