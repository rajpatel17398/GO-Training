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

	r.HandleFunc("/emp", GetEmps).Methods("GET")
	r.HandleFunc("/emp/{id}", GetEmp).Methods("GET")
	r.HandleFunc("/emp", CreateEmp).Methods("POST")
	r.HandleFunc("/emp/{id}", UpdateEmp).Methods("PUT")
	r.HandleFunc("/emp/{id}", DeleteEmp).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))

}

var DB *gorm.DB                                                                            // database variable
var err error                                                                              // variable for error
const DNS = "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local" // constant variable for database url

type Emp struct {
	gorm.Model
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Companyname string `json:"companyname"`
}

func IntiMigration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{}) // connect the database
	if err != nil {                                      // raise the panic for the error
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Emp{}) // enable the auto migrate

}

func GetEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jason")
	var emps []Emp
	DB.Find(&emps)
	json.NewEncoder(w).Encode(emps)

}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var emp Emp
	DB.First(&emp, params["id"])
	json.NewEncoder(w).Encode(emp)

}
func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var emp Emp
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}
func UpdateEmp(w http.ResponseWriter, r *http.Request) { //mix of get and create

	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var emp Emp
	DB.First(&emp, params["id"])         // get the data with database
	json.NewDecoder(r.Body).Decode(&emp) //decode the data of body
	DB.Save(&emp)
	json.NewEncoder(w).Encode(emp)
}
func DeleteEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/jason")
	params := mux.Vars(r) // in this mux we are passing the request and getting the params
	var emp Emp
	DB.Delete(&emp, params["id"])
	json.NewEncoder(w).Encode("the user is deleted")

	//this will soft delete means you can see in database but can't fetch the data from database

}

func main() {
	IntiMigration() // init the migration function
	Initrouter()    // init the routers
}
