package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to slices in GoLang...")

	var fruitList = []string{"Apple", "tomato", "Peach"}
	fmt.Printf("fruitlist type is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruilist is", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	// highScores = append(highScores, 555, 666, 321)

	// fmt.Println(highScores)
	
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slics based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)


}