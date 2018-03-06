/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/19/2018
	Desc:	This file defines all of the api structures
*/

package api

import jwt "github.com/dgrijalva/jwt-go"

// User structure - defines the user objects
type User struct {
	ID       string    `json:"id"`
	Fname    string    `json:"fname"`
	Lname    string    `json:"lname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	LockNum  string    `json:"lockNum"`
	KeyNum   string    `json:"keyNum"`
	Security *Security `json:"security"`
}

// UserLoginState structure - structure passed when a use logs in
type UserLoginState struct {
	Email      string `json:"email"`
	LoginState string `json:"loginState"`
	Locks      string `json:"locks"`
}

// UserLoginCheck structure - structure passed when a use logs in
type UserLoginCheck struct {
	Email      string `json:"email"`
	LoginState string `json:"LoginState"`
	Secret     string `json:"secret"`
}

// UserLoginSuccess structure - structure passed when a user loggon is successful
type UserLoginSuccess struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	LoginState string `json:"loginState"`
	Token      string `json:"token"`
}

// UserRegisterStart structure -
type UserRegisterStart struct {
	Email    string `json:"email"`
	TempPass string `json:"tempPass"`
}

// UserRegisterCont structure -
type UserRegisterCont struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Fname      string `json:"fname"`
	Lname      string `json:"Lname"`
	SecurityLv string `json:"SecurityLv"`
}

// JwtClaims - for issue a JWT for authentication
type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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

// Lock structure - attaches directly to the user via UserID User
type Lock struct {
	ID       string `json:"id"`
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
