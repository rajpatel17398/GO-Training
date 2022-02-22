package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is file in golang")
	content := "this is the content string which we are using in this code"
	file, err := os.Create("./mytextfile.txt") // we use os.Create() for create the file

	if err != nil {
		panic(err) //panic() just shut down the program and print the error
	}
	length, err := io.WriteString(file, content) //for getting the lenght of the file we use io.WriteString() wiht filename, and content which we have to put inside
	if err != nil {
		panic(err) //panic() just shut down the program and print the error
	}
	fmt.Println("length is:", length)
	defer file.Close() // always close the file after reding the file
	readFile("./mytextfile.txt")
}

func readFile(filename string) { // make another function to read the file
	databyte, err := ioutil.ReadFile(filename) // for reading the data use ioutil.Readfile(filename)
	checkNilErr(err)                           // just use the function for check the error
	fmt.Println("Text data inside the file is \n", string(databyte))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err) //panic() just shut down the program and print the error
	}
}
