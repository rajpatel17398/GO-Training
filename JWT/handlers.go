package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key") // this is the jwt key, we are using it to sign jwt token

// create the map and create username and password there
var users = map[string]string{

	"user1": "password1",
	"user2": "password2",
}

// create the struct for username and password
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//second struct is claims user for payload

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

}
func Home(w http.ResponseWriter, r *http.Request) {

}
func Refresh(w http.ResponseWriter, r *http.Request) {

}
