package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	var fruitList [4]string //declair the array
	fruitList[0] = "apple"  //initialize the values
	fruitList[1] = "tomato"
	fruitList[3] = "grapes"
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("fruit list is ", len(fruitList))

	var vagiesList = [3]string{"potato", "bhindo", "bataka"} //declare + initialize the array
	fmt.Println("vaggie list", vagiesList)

	var vagiesList1 = [...]string{"potato", "bhindo", "bataka", "bataka"} //declare + initialize the array for unlimited values
	fmt.Println("new vaggie list", vagiesList1)
	fmt.Println("new vaggie list length", len(vagiesList1))

	// fruitList[0]=""
	// fruitList[0]=""

}
