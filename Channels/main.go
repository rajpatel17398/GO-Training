package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("example 'worker 1 start the job 1','worker 1 finished the job 1'")
	channal := make(chan int, 1)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerpool(i, channal)
		wg.Wait()
	}
	for j := 1; j <= 5; j++ {
		channal <- j
	}
}

func workerpool(id int, cha chan int) {
	fmt.Println("worker", id, "start job", cha)
	time.Sleep(time.Second)
	fmt.Println("worker", id, "complete job", cha)
}
