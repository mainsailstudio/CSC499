/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/19/2018
	Desc:	This file defines all of the api structures
*/

package api

// User structure - defines the user objects
type User struct {
	ID       string    `json:"id"`
	Fname    string    `json:"fname"`
	Lname    string    `json:"lname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	LockNum  string    `json:"lockNum"`
	KeyNum   string    `json:"keyNum"`
	TempPass string    `json:"tempPass"`
	Security *Security `json:"security"`
}

// Security structure - the security levels
// Not attached to the user like the others, user attaches to this
type Security struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

// Auth structure - each individual authentication word
// Attaches directly to the user via UserID User
type Auth struct {
	UserID string `json:"userid"`
	Auth   string `json:"auth"`
	Salt   string `json:"salt"`
}

// Lock structure - each key that attaches directly to an auth via authID
// Attaches directly to the user via UserID User
type Lock struct {
	AuthID   string `json:"id"`
	UserID   string `json:"userid"`
	Lock     string `json:"locksAre"`
	LockType string `json:"lockType"`
}

// KeyDisplay structure - what is allowed to be displayed
// Attaches directly to the user via UserID User
type KeyDisplay struct {
	ID      string `json:"id"`
	UserID  string `json:"userid"` // anonymous field for userID
	Key     string `json:"keysAre"`
	KeyType string `json:"keyType"`
}
