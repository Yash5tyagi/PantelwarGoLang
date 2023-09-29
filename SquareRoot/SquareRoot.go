package main

import (
	"fmt"
)

func Sqrt(x float64) {
	// ans := float64(0)
	no := float64(1)
	point := float64(1)
	for current := 0; current < 12; current++ {

		for ; no*no <= x; no += point {

		}

		no = no - point
		// fmt.Println(no)
		point = point / 10
	}

	fmt.Println(no)
}

func main() {
	Sqrt(99)
}

// package main

// import "fmt"

// func main() {
// 	point:=float64(1)
// 	point=point/10
// 	fmt.Println(point)
// }