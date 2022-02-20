package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "welcome here"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our food ")

	// comma ok || comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)
	fmt.Printf("types of the input %T \n", input)

}
