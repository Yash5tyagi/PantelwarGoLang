package main

import (
	"fmt"
	"time"
)

func test(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func test2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go test("Hello")
	test2("World")
}
