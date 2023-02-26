package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in goLang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("FruitsList is ", fruitList)
	fmt.Println("FruitsList Length is ", len(fruitList))


	var vegList = [3] string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is", vegList)
}