/*
	Title:	Authentication package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	This is the package that takes in
*/

package dynauthcore

import (
	"bytes"
	"database/sql"
	dbinfo "dbinfo"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//Authenticate - to do the authentication I suppose.
func Authenticate(locks string, otp string, userid string, iterations int) {
	// first prep auth for comparison
	toHash := locks + otp
	fmt.Println("toHash is", toHash)
	hashedOTP := hashOTP(toHash, iterations) // hash the prepped otp
	fmt.Println("Hashed otp is", hashedOTP)
	if compareAuth(hashedOTP, userid) == true {
		fmt.Println("AUTHENTICATED")
	} else {
		fmt.Println("NO MATCH FOUND")
	}
}

func hashOTP(otp string, iterations int) string {
	hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(otp), iterations)
	if err != nil {
		panic(err)
	}
	hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
	return hashedToString
}

func compareAuth(toCompare string, userid string) bool {
	authSlice := getAuths(userid) // get all of the auths into a slice
	var authenticated bool
	fmt.Println("Auth slice is", authSlice)
	for i := range authSlice {
		fmt.Println("Compare number", i)
		authenticated = false
		err := bcrypt.CompareHashAndPassword([]byte(authSlice[i]), []byte(toCompare))
		fmt.Println("Err is", err)
		if err == nil {
			authenticated = true
			break
		}
	}
	return authenticated
}

func getAuths(userid string) []string {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()
	authSlice := []string{}
	query := "SELECT auth FROM auth" + userid
	locks, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer locks.Close()
	for locks.Next() {
		var auth string
		err := locks.Scan(&auth)
		if err != nil {
			log.Fatal(err)
		}
		authSlice = append(authSlice, auth)
	}
	err = locks.Err()
	if err != nil {
		log.Fatal(err)
	}
	return authSlice
}
