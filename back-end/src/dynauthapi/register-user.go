/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package api

import (
	"bytes"
	"database/sql"
	dbinfo "dbinfo"
	dynauthconst "dynauthconst"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser - create and insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Fname != "" || user.Lname != "" || user.Phone != "" || user.LockNum != "" || user.KeyNum != "" || user.Security != nil {
		fmt.Println("Would update user to match here")
		//createFullUser(user.Fname, user.Lname, user.Phone, user.LockNum, user.KeyNum,)
	} else if user.Email != "" && user.TempPass != "" {
		if checkUserExists(user.Email) == false {
			registerUser(user.Email, user.TempPass)
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, "This email already exists, please use a different email", 400)
		}
	} else {
		http.Error(w, "There was an error with the register user api call, the fields did not match any method", 400)
	}
}

func checkUserExists(email string) bool {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	exists := false
	// search to make sure this email doesn't already exist
	var whoCares string
	row := db.QueryRow("SELECT email FROM users where email = ?", email).Scan(whoCares)
	switch row {
	case sql.ErrNoRows:
		fmt.Println("No rows selected")
		exists = false
	default:
		exists = true
	}

	return exists
}

// registerUse - only takes email and tempPass for simple registration
func registerUser(email string, tempPass string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user is created
	initUser := "INSERT INTO users (id, email) VALUES (DEFAULT, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(email)
	if err != nil {
		panic(err.Error())
	}

	// select the userid of the user that was just created
	// nice example of a simple single row query
	var userid string
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&userid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User id selected was", userid)

	// This is where each unique user password is created
	initUserPass := "INSERT INTO tempPass (userid, pass) VALUES (?, ?)"
	stmtInsPass, err := db.Prepare(initUserPass)
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsPass.Close()

	// create hashed password
	hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(tempPass), dynauthconst.BcryptIterations)
	if err != nil {
		panic(err)
	}
	tempPass = bytes.NewBuffer(hashedPasswordBcrypt).String()

	fmt.Println("Hashed temp pass is", tempPass)
	_, err = stmtInsPass.Exec(userid, tempPass)
	if err != nil {
		panic(err.Error())
	}
}

func createFullUser() {

}
