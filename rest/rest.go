package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initrouter() {
	r := mux.NewRouter()

	r.HandleFunc("/book", GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", GetBook).Methods("GET")
	r.HandleFunc("/book", CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))

}

var DB *gorm.DB                                                                            // database variable
var err error                                                                              // variable for error
const DNS = "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local" // constant variable for database url

type Book struct {
	gorm.Model
	Id         int    `json:"id"`
	Isbn       int    `json:"isbn"`
	Title      string `json:"title"`
	Authorname string `json:"authorname"`
}

func IntiMigration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{}) // connect the database
	if err != nil {                                      // raise the panic for the error
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Book{}) // enable the auto migrate

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jason")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var book Book
	DB.First(&book, params["id"])
	json.NewEncoder(w).Encode(book)

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	json.NewEncoder(w).Encode(book)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) { //mix of get and create

	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var book Book
	DB.First(&book, params["id"])         // get the data with database
	json.NewDecoder(r.Body).Decode(&book) //decode the data of body
	DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var book Book
	DB.Delete(&book, params["id"])
	json.NewEncoder(w).Encode("the user is deleted")

	//this will soft delete means you can see in database but can't fetch the data from database

}

func main() {
	IntiMigration() // init the migration function
	Initrouter()    // init the routers
}
