package main

import "fmt"

func main() {
	fmt.Println("welcome to mypointers")

	var ptr *int // pointer initialization
	fmt.Println("value of pointer is ", ptr)

	Mynumber := 23

	var ptr1 = &Mynumber //make pointer which referncing a value

	fmt.Println("value of actual pointer is ", ptr1)  // pointer value which is address of Mynumber
	fmt.Println("value of actual pointer is ", *ptr1) // adding * before pointer gives the original value

	*ptr1 = *ptr1 * 2 // changing value of pointer which directly change into original Mynumber value
	fmt.Println("new value of pointer", *ptr1)
	fmt.Println("new value ", Mynumber)

}
