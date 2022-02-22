package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in go lang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and the value is %v \n", index, day)
	// }
	value := 1
	for value < 100 {

		if value == 10 {
			goto gotostatement // for using goto statement just use goto keyword then provide the goto statement
		}

		if value == 20 {
			value++
			continue //continue just skip one phase but breaks terminate the whole thing
			// break
		}
		fmt.Println("values are", value)
		value++
	}

gotostatement: // creating goto statement , just add : after any keyword which is not predefined
	fmt.Println("go and buy the chocolate")

}
