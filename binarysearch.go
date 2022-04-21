// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr = []int{5, 5, 5, 6, 7, 7, 8}
// 	// fmt.Println(len(arr))
// 	x := 8
// 	// sort.Ints(arr)

// 	fmt.Println(Binary(arr, x))

// }
// func Binary(array []int, x int) int {

// 	low := 0
// 	high := len(array) - 1
// 	oriarr := []int{}
// 	duparr := []int{}
// 	for low <= high {
// 		mediun := (low + high) / 2

// 		if array[mediun] == x {
// 			for i, j := range array {

// 			}
// 			return mediun + 1
// 		} else if x < array[mediun] {
// 			high = mediun - 1
// 		} else {
// 			low = mediun + 1
// 		}
// 	}
// 	fmt.Println("Something went wrong!!!")
// 	return 0

// }
// package main

// import "fmt"

// func main() {
// 	fmt.Println("finding duplicated using binary search algo")
// 	slice := []int{1, 2, 3, 4, 4, 4, 5, 6, 6, 6, 7}
// 	// x := 5
// 	// fmt.Println(slice)
// 	for i := 0; i < len(slice); i++ {
// 		fmt.Println(slice[i])
// 		target := slice[i]
// 		low := 0
// 		high := len(slice) - 1
// 		// result := []int{}

// 		for low <= high {
// 			mid := (low + high) / 2
// 			if slice[mid] < target {

// 				low = mid + 1
// 			} else if slice[mid] > target {

// 				high = mid - 1
// 			} else {
// 				fmt.Println("true", mid+1)

// 			}

// 		}

// 	}
// }
package main

import "fmt"

func findDuplicate(nums []int) int {
	low := 1
	high := len(nums) - 1
	for low <= high {
		mid := (low + (high-low)/2)
		cnt := 0
		fmt.Println("mid", mid)
		for _, a := range nums {
			if a <= mid {
				cnt++
			}
			fmt.Println("cnt", cnt)

		}
		if cnt <= mid {
			low = mid + 1

		} else {
			high = mid - 1
		}

	}

	return low
}

func main() {
	fmt.Println(findDuplicate([]int{2, 6, 6, 6, 6, 1}))
}
