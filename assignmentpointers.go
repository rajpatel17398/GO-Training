package main

import "fmt"

func main() {

	// Double pointer & nil pointer
	var a *int // pointers
	var b **int
	var c *int // empty pointer

	val := 10

	a = &val // storing the address of val in pointer a
	b = &a   // storing the address of the a pointer to pointer b

	fmt.Printf("Value of a = %d\n", val)
	fmt.Printf("Value available at *ptr = %d\n", *a) // using dereferencing pointer get the values
	fmt.Printf("Value available at **pptr = %d\n", **b)
	fmt.Printf("The value of ptr is : %x\n", c) // nil pointer

	// Panic function

	if a := 50; a > 49 {
		panic("the age is greater than 49")
	}

	// Pointer arithmetic is not possible in GO lang like C language.

	val1 := 10
	val2 := &val1
	val2 = val2 + 1
	fmt.Println(val2)

}
