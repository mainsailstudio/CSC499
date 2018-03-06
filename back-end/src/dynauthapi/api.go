/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/gorilla/handlers"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni" // to implement the JWT middleware
	// JWT middleware
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Init hashs var as a slice Hash struct
var users []User

// Get all books
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get all books
func getUsersReal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sqlString := "select * from users"
	usersDB, err := GetJSONFromSQL(sqlString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// json.NewEncoder(w).Encode(usersDB)
	fmt.Println(usersDB)
	fmt.Fprintf(w, usersDB) // prints to browser
}

// // Get single book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Gets params
// 	// Loop through books and find one with the id from the params
// 	for _, item := range auths {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Update book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// }

// // Delete book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(books)
// 	}
// }

// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range tokens {
// 		if item.ID == params["tokens"] {

// 			break
// 		}
// 		json.NewEncoder(w).Encode(books)
// 	}
// }

// just to test quickly
func testAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Only view this if you are authenticated")
}

// StartAPI - Start the api for testing
func StartAPI() {

	// Hardcoded data - @todo: add database
	testUser := User{ID: "1", Fname: "Connor", Lname: "Peters", Email: "design@mainsailstudio.com", Phone: "5854782678",
		Security: &Security{
			ID:   "1",
			Name: "Security level 1",
			Desc: "This is security level 1"}}
	userLocks := Lock{ID: "1", UserID: testUser.ID, Lock: "1", LockType: "1"}
	_ = userLocks
	userAuths := Auth{UserID: testUser.ID, Auth: "testauthIguess", Salt: "Iguess"}
	_ = userAuths
	userKeyDisplay := KeyDisplay{UserID: testUser.ID, Key: "test1", KeyType: "1"}
	_ = userKeyDisplay

	// Route handles & endpoints
	// Init router
	fmt.Println("Starting mux")

	// read private key and use that as the secret
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(pkSecret))

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
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Allow-Methods", "GET")
		// w.Header().Set("Access-Control-Allow-Methods", "POST")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users-real", getUsersReal).Methods("GET")
	r.HandleFunc("/register", CreateUser).Methods("POST")
	r.HandleFunc("/login-start", GetLoginState).Methods("POST")
	r.HandleFunc("/login-finish", LoginUser).Methods("POST")

	// testing authentication
	// r.HandleFunc("/api", GetAPIBase).Methods("GET")
	r.HandleFunc("/api/users", GetAPIUsers).Methods("GET")
	r.HandleFunc("/api/authenticate", PostAPIAuth).Methods("GET")
	//r.HandleFunc("/register", CreateUser).Methods("POST")

	r.Handle("/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(testAuth)),
	))

	r.Handle("/register-continue", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(CreateUserContinue)),
	))

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)

}
