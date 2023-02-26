package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in goLang...")

	languages := make(map[string]string)

	languages["JS"] = "javaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languages ", languages)
	fmt.Println("Js short form of", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages ", languages)

	// Loops are interesting in goLang

	for key, value := range languages{
		fmt.Printf("For Key %v, value is %v \n", key, value)
	}





}