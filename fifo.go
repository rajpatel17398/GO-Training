package main

import "fmt"

func sqare(c, quit chan int) {
	var num int

	for {

		select {
		case c <- num:
			num = num * num

			// fmt.Println(num)

		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	c := make(chan int, 10)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	sqare(c, quit)

}
