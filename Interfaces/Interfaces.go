package main

import "fmt"

type mygeneric interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func main() {
	var my mygeneric
	var another mygeneric

	c := circle{10}
	r := rectangle{10, 20}

	my = c
	fmt.Println(my.area())

	another = r
	fmt.Println(another.area())
}
