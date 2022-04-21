package main

import "fmt"

func main() {
	// count of duplicate numbers
	var arr = [...]int{1, 2, 3, 4, 5, 6, 1, 2, 1, 2}
	fmt.Println(arr)

	var mapp = map[int]bool{}
	result := []int{}
	result1 := []int{}
	for i := 0; i < len(arr); i++ {
		if mapp[arr[i]] == false {
			mapp[arr[i]] = true
			result = append(result, arr[i])
			fmt.Println("Unique Values are:", result)
		} else if mapp[arr[i]] == true {
			result1 = append(result1, arr[i])
			fmt.Println("Duplicates values are:", result1)
			fmt.Println("Duplicates value's count are:", len(result1))

		}
	}
}
