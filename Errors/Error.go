package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string { //Jab bhi error print hoga yhe automatically chalega
	return "cannot Sqrt negative number: -2"
}

func Sqrt(x float64) (float64, error) {
	var a ErrNegativeSqrt = 2.4
	if x < 0 {
		a = ErrNegativeSqrt(x)

		return x, a
	} else {
		// fmt.Println(a)
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

		return no, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) //no,error
}
