package main

import "fmt"

func main() {

	// defer always runs last and if you have all defer statement then it is lifo stack
	defer fmt.Println("this is defer statement")
	fmt.Println("this is the defer")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

}
