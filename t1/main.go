package main

import (
	"fmt"
	"time"
	appKafka "vs_code/t1/kafka"
)

func main() {
	fmt.Println("this is the task1 given by Chekri Sir")
	go appKafka.StartKafka()

	fmt.Println("kafka has been started")

	time.Sleep(10 * time.Minute)

}
