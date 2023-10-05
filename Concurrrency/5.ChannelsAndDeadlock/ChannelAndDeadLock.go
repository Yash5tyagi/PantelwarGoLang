package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning Channels")

	wg := &sync.WaitGroup{}
	                    
	                 //Buffer value
	myCh := make(chan int, 1) //Channel will transfer int value
	// myCh <- 5              //I cannot use a single channel,I can only initialize it or giving it a value when it communicates with another channel
	// fmt.Println(<-myCh)
	wg.Add(2)

	// Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-ch

		fmt.Println(val)
		fmt.Println(isChanOpen)
		// fmt.Println(<-ch) //Listening a channel value is allowed even after closing the channel
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	// Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		close(ch) //Cannot assign value to channel after this
		// ch <- 5
		// ch <- 6
		// ch <- 7
		// close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
