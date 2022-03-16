package main

import (
	"log"
	"net/http"
)

func main() {
	// create the handlers
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	// now we have to start the function
	log.Fatal(http.ListenAndServe(":8080", nil)) // we are not using any router framework so we write nil here.

}
