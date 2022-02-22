package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano()) // we will get the random number every time
	diceNumber := rand.Intn(6) + 1
	fmt.Println(" DICE NUMBER IS ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can move 1 steps")
	case 2:
		fmt.Println("you can move 2 steps")
	case 3:
		fmt.Println("you can move 3 steps")
		fallthrough // fallthrough executes the same case and push the compiler to the next statement so that is also execute with current statement.
	case 4:
		fmt.Println("you can move 4 steps")
		fallthrough

	case 5:
		fmt.Println("you can move 5 steps")
	case 6:
		fmt.Println("you can move 6 steps and you can open at first or you can re-roll")
	default:
		fmt.Println("this is default value")

	}

}
