package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// This format number needs to be same
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 16, 01, 01, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Mon"))

}