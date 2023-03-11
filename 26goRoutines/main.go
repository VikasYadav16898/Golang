package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup  //usually these are pointers
var mut sync.Mutex //usually these are made pointers

func main() {
	// go greeter("Hello")
	// greeter("world")


	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)



}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	 defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint...")
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	
	fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
}