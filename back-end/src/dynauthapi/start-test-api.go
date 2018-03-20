/*
	Title:	Serve a REST-like API for a usability test using Mux
	Author:	Connor Peters
	Date:	3/17/2018
	Desc:	This is the test controller of the application that presents an api over TLS
	NOTES:	For testing purposes, the negroni middleware has to be subbed out for a normal handlefunc to prevent CORS errors
*/

package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartTestAPI - Start the TEST api for mux api and serve it over HTTP (eventually will be TLS once implemented)
func StartTestAPI() {

	// Init router
	fmt.Println("Starting mux")
	r := mux.NewRouter()

	// // Read private key and use that as the secret
	// pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Declare a jwtMiddleware variable that is used with the negroni JWT middleware
	// // This checks to make sure each JWT token is validated using the private key read above
	// jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		return pkSecret, nil
	// 	},
	// 	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// 	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// 	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// 	SigningMethod: jwt.SigningMethodHS256,
	// })

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Basic handles for API
	r.HandleFunc("/test/login-start", testLoginLevel).Methods("POST")
	r.HandleFunc("/test/login-finish", testLogin).Methods("POST")
	r.HandleFunc("/test/register", testRegister).Methods("POST")
	r.HandleFunc("/test/log-config", logConfigActivity).Methods("POST")
	r.HandleFunc("/test/log-login", logLoginActivity).Methods("POST")
	r.HandleFunc("/test/register-auth", testRegisterAuth).Methods("POST")
	r.HandleFunc("/test/register-keys", testRegisterKeys).Methods("POST")
	r.HandleFunc("/test/register-pass", testRegisterPass).Methods("POST")
	r.HandleFunc("/test/get-keys", testGetUserKeys).Methods("GET")

	// // restricted API call to register a user with a password
	// // requires a proper JWT token to access
	// r.Handle("/test/register-pass", negroni.New(
	// 	negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(testRegisterPass)),
	// ))

	// // restricted API call to register a user's auths
	// // requires a proper JWT token to access
	// r.Handle("/test/register-auth", negroni.New(
	// 	negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(testRegisterAuth)),
	// ))

	// Run the handler through CORS for testing
	handler := cors.Default().Handler(r)

	// Listen and serve the API over HTTP
	// TODO: have this serve over TLS for the server
	http.ListenAndServe(":8080", handler)

	// Listen and serve the API over TLS (HTTPS)
	// err := http.ListenAndServeTLS(":443", "server.crt", "server.key", handler)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
