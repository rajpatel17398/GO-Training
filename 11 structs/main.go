package main

import "fmt"

//structs are the alternative of class methods.

func main() {
	fmt.Println("welcome to the struct")

	//there is no inheritance in golang, no super or parent

	raj := User{"Raj", "raj@gmail.com", true, 23}

	fmt.Println(raj)

	fmt.Printf("raj details are: %+v \n", raj)                    //you should write %+v for all details
	fmt.Printf("name is %v and email is %v", raj.Name, raj.Email) // you should write %v for perticular detail
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
