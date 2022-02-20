package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study")

	presenttime := time.Now()

	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 Monday 15:04:05"))

	createddate := time.Date(2020, time.March, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createddate)

	fmt.Println(createddate.Format("01-02-2006 Monday"))
}
