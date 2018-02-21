/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"net/http"

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

// StartAPI - Start the api for testing
func StartAPIPlease() {

	// Hardcoded data - @todo: add database
	users = append(users, User{ID: 1, Fname: "Connor", Lname: "Peters", Email: "design@mainsailstudio.com", Phone: "5854782678",
		Security: &Security{
			ID:   1,
			Name: "Security level 1",
			Desc: "This is security level 1"},
		Locks: &Lock{
			ID:       1,
			UserID:   1,
			Lock:     "1",
			LockType: 1},
		Auths: &Auth{
			Auth: "testauthIguess",
			Salt: "Iguess"},
		KeysDisplay: &KeyDisplay{
			ID:      1,
			UserID:  1,
			Key:     "test1",
			KeyType: 1}})

	// Route handles & endpoints
	// Init router
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

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")

	// // Start server
	// log.Fatal(http.ListenAndServe(":8000", r))
	// Start server
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)

}
