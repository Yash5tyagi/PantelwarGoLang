package main

import "fmt"

func main() {
	a:=10
	temp:=&a
	fmt.Println(*temp)

	*temp=20
	fmt.Println(*temp)
}