package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	// var wg sync.WaitGroup //pointer
	wg := &sync.WaitGroup{}
	// var mut sync.Mutex    //pointer
	mut := &sync.Mutex{}  //We use it avoid race condition

	var score = []int{0}

	fmt.Println("<------Race Condition------>")

	wg.Add(3) // or wg.Add(1)  //Just assume it as a stack that pushes this task that we should wait for 1 task
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score) //We cannot predict ki kaun sa thread phele chal jaye because these are just independent kind of threads
}

//Now to check whether my program has race condition or not
// go run --race
