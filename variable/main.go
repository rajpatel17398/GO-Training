package main

import "fmt"

const Pub = 1000000

func main() {
	fmt.Println("variable")
	var username string = "this is my main"
	fmt.Println(username)
	fmt.Printf("variable is of type %T", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type %T \n", isLoggedin)
	fmt.Println(Pub)
	fmt.Printf("variable is of type %T \n", Pub)

}
