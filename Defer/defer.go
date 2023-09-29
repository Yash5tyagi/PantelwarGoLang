package main

import "fmt"

func add(x int){
fmt.Println(x)
}

func main() {
    val:=10
	
	defer add(val)
	val=1
	fmt.Println("hello")
	
	
	fmt.Println("hello")
	
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
}