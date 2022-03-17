package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key") // this is the jwt key, we are using it to sign jwt token

// create the map and create username and password there
var users = map[string]string{

	"user1": "password1",
	"user2": "password2",
}

// create the struct for username and password
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//second struct is claims user for payload

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//if the username and password from body we are getting is matching in map's username and the password or
// you can say instead of map the data base then the claim will create the jwt tokens

func Login(w http.ResponseWriter, r *http.Request) {

	//1 we are getting the credentials username and password from body

	var credentials Credentials

	//2 it will decode the message
	err := json.NewDecoder(r.Body).Decode(&credentials)

	//3 if any the credentials are not matching then it will throw the error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//4 we are checking if the data is available or not
	expectedPassword, ok := users[credentials.Username] //storing the usename to the expectedpassword and ok

	//5 if the data is not available or the passeword is not matching that we have that means this is unauthorize access
	if !ok || expectedPassword != credentials.Password { // if the password is not match then it will state as unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//6 once all are good then we are creating the expiration time for out token
	// creating expiration time
	expirationTime := time.Now().Add(time.Minute * 5)
	//7 then we are creating object of the claims that we created before and now we are giving the data
	claims := Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//8 we are create the token and pass the algorithm SigningMethodHS256
	//create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//9 create the token string through passing the jwtkey
	// create the token string
	tokenString, err := token.SignedString(jwtKey)

	//10 if there are the error we are throw the internal server error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//11 we are set the info of token tokenstring and expiration time in cookie
	//creating the cookie

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

//in home method what ever token we are getting back are the exactly the same which we are sending back, if it is
//authontication if the token is valid then home will allow user to let on home page
func Home(w http.ResponseWriter, r *http.Request) {
	//1 geting the cookie and error
	cookie, err := r.Cookie("token")

	//2 checking the error types
	if err != nil {
		//3 if the error is nocookie the there is unauthorized
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
		//4 if it cookie is there but still error then it is bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//5 if it is valid then we are getting token sting which is cookie value
	tokenStr := cookie.Value
	//6 creating the reference of Claims struct
	claims := &Claims{}

	//7 with jwt.parsewithclaims we are passsing the tokenstr and calims and the function which is returning
	//the interface

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {

		//8 passing the jwtkey over here and this will return the token tkn
		return jwtKey, nil

	})

	// 9 again we are checking the error here

	if err != nil {
		//10 if the error is signature invalid then it will return the
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//11 if error is other type so we are writing bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//12 if there is no error and token is valid then we are sending out data back to the client
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	//sending the data back to the client
	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

}

// we have cookie now but it will expire on some point for reuse the cookie we are making the refresh function
func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	//2 checking the error types
	if err != nil {
		//3 if the error is nocookie the there is unauthorized
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
		//4 if it cookie is there but still error then it is bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//5 if it is valid then we are getting token sting which is cookie value
	tokenStr := cookie.Value
	//6 creating the reference of Claims struct
	claims := &Claims{}

	//7 with jwt.parsewithclaims we are passsing the tokenstr and calims and the function which is returning
	//the interface

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {

		//8 passing the jwtkey over here and this will return the token tkn
		return jwtKey, nil

	})

	// 9 again we are checking the error here

	if err != nil {
		//10 if the error is signature invalid then it will return the
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//11 if error is other type so we are writing bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//12 if there is no error and token is valid then we are sending out data back to the client
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	//13 if the token is about to expire then we will create the new token
	// if time.Unix(claims.ExpiresAt,0).Sub(time.Now())>30*time.Second{
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// 14 we have to create tha new token
	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime.Unix()

	//8 we are create the token and pass the algorithm SigningMethodHS256
	//create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//9 create the token string through passing the jwtkey
	// create the token string
	tokenString, err := token.SignedString(jwtKey)

	//10 if there are the error we are throw the internal server error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshed_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
