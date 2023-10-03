package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
	a := 0
	b := 1

	ans := func(i int) int {
		if i == 0 || i == 1 {
			return i
		} else {
			sum := a + b
			a = b
			b = sum
			return sum
		}
	}

	return ans
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
