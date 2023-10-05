package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	websiteList := []string{
		"https://www.google.com",
		"https://go.dev",
	}

	for i := 0; i < len(websiteList); i++ {
		go getStatusCode(websiteList[i])
		wg.Add(1) //Syntax
	}

	// time.Sleep(1 * time.Second)
	wg.Wait()
}

func getStatusCode(s string) {
	defer wg.Done()
	res, err := http.Get(s)
	if err != nil {
		fmt.Println("Error Agaya Bhai")
	} else {
		mut.Lock()
		signals = append(signals, s)
		mut.Unlock()

		fmt.Println(res.Status)
		fmt.Println(signals)
	}
}
