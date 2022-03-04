package main

import (
	"fmt"
)

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {
	duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}

	// printslice(duplicate)
	fmt.Println("Total Words are", duplicate)

	dup_map := dup_count(duplicate)

	//fmt.Println(dup_map)

	for k, v := range dup_map {
		fmt.Printf(" %s , Count : %d\n", k, v)
	}

}
