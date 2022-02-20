package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the slices")

	var fruitList = []string{"apple", "tometos", "peach"} // this is slice
	fmt.Printf("what type of data type of fruitlist is %T \n", fruitList)
	fruitList = append(fruitList, "mango", "banana") // this is how we add values into slice with append method
	fmt.Println(fruitList)

	fruitList = append(fruitList[:4]) // this is how we slice the array or SLICE
	fmt.Println(fruitList)

	highscores := make([]int, 4) // make slice from make function
	highscores[0] = 0            // add the values according to index
	highscores[1] = 10
	highscores[2] = 45
	highscores[3] = 80
	// highscores[4] = 81

	highscores = append(highscores, 77, 6, 777) // since the index is full and there is no memory available for new value that's why use the append method, which reallocate the memory

	fmt.Println("highscores are ", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores) // this is how we sort any slice
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// how to remove the value from the slice
	var cources = []string{"java", "GO", "python", "angular"}
	fmt.Println(cources)
	var index int = 0
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
