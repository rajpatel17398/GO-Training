package main

import (
	"fmt"
	"vs_code/testP2"
	"vs_code/testPackage"

	"github.com/uniplaces/carbon"
)

func main() {
	// fmt.Println("Hellooo")
	fmt.Println(testPackage.Cal(2, 3))
	fmt.Println(testP2.Sub(5, 1))

	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	fmt.Println("this changes happened before second commit")
	fmt.Println("this changes happened on github")
	fmt.Println("this changes happened before pushing")

}
