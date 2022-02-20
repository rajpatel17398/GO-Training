package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string) // we can also make maps from make function just write map then in [] declare key type and then after than provide value type

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("py key of:", languages["py"])

	//for delete from the map of slice we can use delete function then add key inside the map

	delete(languages, "js") // add map , key of value which you want to delete
	fmt.Println("List of all languages:", languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("for key %v, vakue is %v \n", key, value)
	}

}
