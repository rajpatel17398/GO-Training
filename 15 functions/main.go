package main

import "fmt"

func main() {
	result := adder(3, 5)
	fmt.Println("functions in golang", result)
	greeter2()
	greeter()
	proresult := proadder(2, 5, 7, 8, 9, 01, 10, 20)
	fmt.Println("Pro Result is ", proresult)

}
func adder(i, j int) int {

	return i + j

}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("this is the greeter function")

}
func greeter2() {
	fmt.Println("this is greeter 2")
}
