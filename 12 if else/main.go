package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "intense user"
	} else {
		result = "you are not the user :)"
	}
	fmt.Println(result)

	if 7%2 == 0 {
		fmt.Println("the number is even")

	} else {
		fmt.Println("Number is odd")
	}

	if num := 10; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("number is 10 or greater than 10")
	}

}
