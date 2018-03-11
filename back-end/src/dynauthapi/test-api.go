/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartAPI - Start the api for testing
func StartTestAPI() {

	// Route handles & endpoints
	// Init router
	fmt.Println("Starting mux")

	// read private key and use that as the secret
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return pkSecret, nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	r.HandleFunc("/test/login-token", testGiveToken).Methods("GET")
	r.HandleFunc("/test/register", testRegister).Methods("POST")
	r.HandleFunc("/test", testTestAPI).Methods("GET")

	r.Handle("/test/login", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(testLogin)),
	))

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)

}
