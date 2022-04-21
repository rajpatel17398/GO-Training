package main

import (
	"fmt"
	"sort"
)

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func main() {
	fmt.Println("Welcome to soap factory")
	fmt.Println("Enter the ammount of soup that you want to supply")
	var quantity int
	fmt.Scanln(&quantity)
	fmt.Println("Enter the soap sizes")
	x := input([]int{}, nil)
	sort.Ints(x)

	var slice = []int{}

	for index, _ := range x {
		smallnumber := x[index]
		totalquantity := quantity / smallnumber
		fmt.Printf("you can make %v soaps of %v Oz OR \n", totalquantity, smallnumber)

		slice = append(slice, totalquantity)
		sort.Ints(slice)

		index++

	}
	fmt.Println("the smallest number of soaps that are made with the ammount you added are", slice[0])

}
