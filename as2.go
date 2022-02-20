// package main

// import (
// 	"fmt"
// )

// func main()  {

// 	stri := "just reverse the string"

// 	for _, i := range stri {

// 		defer fmt.Println(string(i))

// 	}

// }
// package main

// import (
// 	"fmt"
// )

// func Lengthmaxsubstring(s string) int {
// 	start, max := 0, 0
// 	if len(s) != 0 {
// 		mapp := make(map[int32]int) //int32 is known as rune
// 		for i, c := range s {       //i is for index, c for character , {a,b,c,d,a,}
// 			if k, ok := mapp[c]; ok && start <= k { //k=0 cause we have same character, then start again from
// 				//b to get unique character
// 				start = k + 1 // start = 1
// 			}
// 			mapp[c] = i          // a=>0, b=>1, c=>2, d=>3 , a=>4
// 			if max < i-start+1 { // max=1,2,3,4
// 				max = i - start + 1
// 			}
// 		}
// 	}
// 	return max
// }
// func main() {
// 	stringg := Lengthmaxsubstring("abcdabcaba")
// 	fmt.Println(stringg)
// }
