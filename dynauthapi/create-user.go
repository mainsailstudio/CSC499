/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"net/http"
)

// CreateUser - create and insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = -1 // not safe
	// user.Security.ID = 2              // not safe
	// user.Locks.ID = 2                 // not safe
	// user.Locks.UserID = user.ID       // not safe
	// user.KeysDisplay.UserID = user.ID // not safe
	json.NewEncoder(w).Encode(user)
}

func insertUser() {
	// make sure to insert user without IDs

}
