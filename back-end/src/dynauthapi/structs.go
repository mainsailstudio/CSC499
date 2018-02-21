/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/19/2018
	Desc:	This file defines all of the api structures
*/

package api

// User structure - defines the user objects
type User struct {
	ID          int         `json:"id"`
	Fname       string      `json:"fname"`
	Lname       string      `json:"lname"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	LockNum     int         `json:"lockNum"`
	KeyNum      int         `json:"keyNum"`
	Security    *Security   `json:"security"`
	Locks       *Lock       `json:"locks"`
	Auths       *Auth       `json:"auths"`
	KeysDisplay *KeyDisplay `json:"keysDisplay"`
}

// Security structure - the security levels
type Security struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

// Auth structure - each individual authentication word
type Auth struct {
	Auth string `json:"auth"`
	Salt string `json:"salt"`
}

// Lock structure - each key that attaches directly to an auth via authID
type Lock struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userid"`
	Lock     string `json:"locksAre"`
	LockType int    `json:"lockType"`
}

// KeyDisplay structure - what is allowed to be displayed
type KeyDisplay struct {
	ID      int    `json:"id"`
	UserID  int    `json:"userid"`
	Key     string `json:"keysAre"`
	KeyType int    `json:"keyType"`
}
