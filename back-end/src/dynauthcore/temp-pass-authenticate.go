/*
	Title:	Temp pass authentication
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:	To allow the user to login via their temp pass
*/

package dynauthcore

import (
	"database/sql"
	dbinfo "dbinfo"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// TempPassAuth - grab the user's temp pass and compare it
func TempPassAuth(userid string, tempPass string) bool {
	authenticated := false

	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// select the temp pass of the user
	var passHash string
	err = db.QueryRow("SELECT pass FROM tempPass WHERE userid = ?", userid).Scan(&passHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pass hash selected", passHash)

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(tempPass))
	fmt.Println("Err is", bcryptErr)
	if bcryptErr == nil {
		authenticated = true
	}

	return authenticated
}
