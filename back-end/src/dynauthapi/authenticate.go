/*
Title:	Registering the user via
	Author:	Connor Peters
	Date:	2/26/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

// GetAPIUsers - to get api users for testing
func GetAPIUsers(w http.ResponseWriter, r *http.Request) {
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

// PostAPIAuth - to get api users for testing
func PostAPIAuth(w http.ResponseWriter, r *http.Request) {
	// going to assume they are authenticated here for testing and going to issue a valid JWT
	token := Token{Token: "fake-jwt-token"}
	jsonToken, _ := json.Marshal(token)
	fmt.Fprintf(w, string(jsonToken)) // prints to browser

}
