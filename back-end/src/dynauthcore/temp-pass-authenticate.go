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
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// TempPassAuth - grab the user's temp pass and compare it
func TempPassAuth(userid string, tempPass string) (bool, error) {
	authenticated := false

	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return false, errors.New("Unable to connect to the database in the TempPassAuth function")
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

	return authenticated, nil
}

// TestPassAuth - grab the user's test pass and compare it
// THIS IS FOR TEST USERS ONLY
func TestPassAuth(userid string, tempPass string) (bool, error) {
	authenticated := false

	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return false, errors.New("Unable to connect to the database in the TestPassAuth function")
	}
	defer db.Close()

	// select the temp pass of the user
	var passHash string
	err = db.QueryRow("SELECT pass FROM testPass WHERE userid = ?", userid).Scan(&passHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pass hash selected", passHash)

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(tempPass))
	fmt.Println("Err is", bcryptErr)
	if bcryptErr == nil {
		authenticated = true
	}

	return authenticated, nil
}
