package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("this is handling urls in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryparams := result.Query() // stores data in key value pair and it is in string format
	fmt.Printf("the types of query params is: %T \n", queryparams)
	fmt.Println(queryparams["coursename"])

	for _, val := range queryparams {
		fmt.Println("params are", val)
	}

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=raj",
	}
	otherURL := partsofURL.String()
	fmt.Println(otherURL)

}
