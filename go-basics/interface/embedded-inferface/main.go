package main

type Sort interface {
	BubbleSort()
}

type List interface { // List has Len() and BubbleSort()
	Sort
	Len()
}
