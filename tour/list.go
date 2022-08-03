package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (this *List[T]) Size() (i int) {
	i = 1
	for current := this; current.next != nil; current, i = current.next, i+1 {
	}
	return i
}

func main() {
	// must specify an actual type when instantiate a generic type
	var strings *List[string] = &List[string]{nil, "abc"}
	strings.next = &List[string]{nil, "xyz"}
	fmt.Printf("%v\n", strings.Size())

	var ints *List[int] = &List[int]{nil, 1}
	ints.next = &List[int]{nil, 2}
	ints.next.next = &List[int]{nil, 3}
	fmt.Printf("%v\n", ints.Size())
}
