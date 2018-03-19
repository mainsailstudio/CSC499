/*
	Title:	Serve a REST-like API using Mux
	Author:	Connor Peters
	Date:	3/15/2018
	Desc:	This is the main controller of the application that presents an api over TLS
	NOTES:	For testing purposes, the negroni middleware has to be subbed out for a normal handlefunc to prevent CORS errors
*/

package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"  // to implement the JWT middleware for JWT tokens
	jwt "github.com/dgrijalva/jwt-go" // JWT middleware
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartAPI - Start the api for mux api and serve it over HTTP (eventually will be TLS once implemented)
func StartAPI() {

	// Init router
	fmt.Println("Starting mux")
	r := mux.NewRouter()

	// Read private key and use that as the secret
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}

	// Declare a jwtMiddleware variable that is used with the negroni JWT middleware
	// This checks to make sure each JWT token is validated using the private key read above
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return pkSecret, nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Allow-Methods", "GET")
		// w.Header().Set("Access-Control-Allow-Methods", "POST")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Basic handles for API
	// r.HandleFunc("/users", getUsers).Methods("GET")
	// r.HandleFunc("/users-real", getUsersReal).Methods("GET")
	r.HandleFunc("/register", CreateUser).Methods("POST")
	r.HandleFunc("/login-start", GetLoginState).Methods("POST")
	r.HandleFunc("/login-finish", LoginUser).Methods("POST")

	// testing authentication
	// r.HandleFunc("/api", GetAPIBase).Methods("GET")
	//r.HandleFunc("/register", CreateUser).Methods("POST")

	// Testing a JWT authenticated function for the API
	r.Handle("/register-continue", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(CreateUserContinue)),
	))

	// Run the handler through CORS for testing
	handler := cors.Default().Handler(r)

	// Listen and serve the API over HTTP
	http.ListenAndServe(":8080", handler)

	// Listen and serve the API over TLS
	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	// if err != nil {
	//     log.Fatal("ListenAndServe: ", err)
	// }
}
