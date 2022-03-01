package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book struct (model)
type Book struct {
	ID     string  `jason:"id"`
	Isbn   string  `jason:"isbn"`
	Title  string  `jason:"Title"`
	Author *Author `jason:"author"`
	// ID string `jason:"id"`
}
type Author struct {
	Firstname string `jason:"firstname"`
	Lastname  string `jason:"lastname"`
}

//init books var as a slice Book struct
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(books)
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r) // get params
	//loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// update the book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}

	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			break
		}

	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	fmt.Println("this is the rest api ")

	//init router
	r := mux.NewRouter()
	// mock data
	books = append(books, Book{ID: "1", Isbn: "448999", Title: "Book one", Author: &Author{Firstname: "Raj", Lastname: "Patel"}})
	books = append(books, Book{ID: "2", Isbn: "449000", Title: "Book two", Author: &Author{Firstname: "Raj2", Lastname: "Patel"}})
	books = append(books, Book{ID: "3", Isbn: "449001", Title: "Book three", Author: &Author{Firstname: "Raj3", Lastname: "Patel"}})

	//route handlers / Endpoint

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
