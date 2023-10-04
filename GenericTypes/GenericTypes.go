package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	n1 := List[int]{nil, 10}

	n2 := List[int]{nil, 20}
	n3 := List[int]{nil, 30}

	n1.next = &n2
	n2.next = &n3

	for i := &n1; i != nil; i = i.next {
		fmt.Println((*i).val)
	}
}
