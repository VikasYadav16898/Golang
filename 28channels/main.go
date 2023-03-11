package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in GoLang...")

	// myCh := make(chan int)
	// next argument decides the number of values acceptable for channels to pass allowed
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// push values in ch unique way, channels allows passing a value only if someone is listening to it 
	// myCh <- 5
	// get value from ch unqiue way
	// fmt.Println(<-myCh)

	wg.Add(2)
	// RECEIVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(myCh)
		wg.Done()
	}(myCh, wg)
	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}