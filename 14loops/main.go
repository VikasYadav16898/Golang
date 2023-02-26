package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GoLang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// // Type 1
	// for d:=0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// // Type 2
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// // Type 3
	// for index, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }


	// Do while loop
	rogueValue := 1
	for rogueValue < 10{

		if rogueValue == 2{
			goto lco
		}

		if rogueValue == 5 {
			break
		}

		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

	lco:
		fmt.Println("Jumping at learn code online")


}