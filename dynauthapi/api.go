/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Auth structure - each individual authentication word
type Auth struct {
	ID   string `json:"id"`
	Auth string `json:"auth"`
	Salt string `json:"salt"`
}

// Key structure - each key that attaches directly to an auth via authID
type Key struct {
	ID     string `json:"id"`
	AuthID string `json:"authid"`
}

// Delim stucture - the delimiter between each key + auth pair
type Delim struct {
	ID    string `json:"id"`
	Delim string `json:"delim"`
}

// Hash structure - the entire hash with keys, delimiters, and auths
type Hash struct {
	HashID string `json:"hashid"`
	Auth   *Auth  `json:"auth"`
	Key    *Key   `json:"key"`
	Delim  *Delim `json:"delim"`
}

// Init hashs var as a slice Hash struct
var hashs []Hash

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hashs)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tokens {
		if item.ID == params["tokens"] {

			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

// StartApi - Start the api plox!
func StartApi() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/auth/{token}", getAuths).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
