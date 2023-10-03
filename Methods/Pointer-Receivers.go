package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func (v *Vertex) Scale(f float64) { //Pointer bhejunga to value update hojayegi
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v)
}

// func (v *Vertex) Scale(f float64) { //Value bhejunga to alag copy banegi,update nahi // hogi value
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// 	fmt.Println(v)
// }

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
