package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("this is file in golang")

	databyte, err := ioutil.ReadFile("./mytextfile.txt")
	checkNilErr(err) // just use the function for check the error
	fmt.Println("Text data inside the file is \n", string(databyte))
	duplicate := []string{string(databyte)}
	fmt.Println(duplicate)
	dup_map := dup_count(duplicate)

	for k, v := range dup_map {
		fmt.Printf(" %s , Count : %d\n", k, v)
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) //panic() just shut down the program and print the error
	}
}
