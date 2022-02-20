// // package main

// // import (
// // 	"fmt"
// // 	"vs_code/testP2"
// // 	"vs_code/testPackage"

// // 	"github.com/uniplaces/carbon"
// // )

// // func main() {
// // 	// fmt.Println("Hellooo")
// // 	fmt.Println(testPackage.Cal(2, 3))
// // 	fmt.Println(testP2.Sub(5, 1))

// // 	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

// // 	fmt.Println("this changes happened before second commit")
// // 	fmt.Println("this changes happened on github")
// // 	fmt.Println("this changes happened before pushing")
// // 	//update the code
// // 	fmt.Println("")
// // 	fmt.Println("")

// // }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	case today + 3:
// 		fmt.Println("In three days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }
