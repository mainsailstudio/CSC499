/*
	Title:	Authentication package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	This is the package that takes in
*/

package dynauthcore

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticateBcrypt - to do the authentication I suppose.
func AuthenticateBcrypt(locks string, otp string, userid string, iterations int) {
	// first prep auth for comparison
	toHash := locks + otp
	fmt.Println("toHash is", toHash)
	// hashedOTP := hashBcrypt(toHash, iterations) // hash the prepped otp
	// fmt.Println("Hashed otp is", hashedOTP)
	authenticated, err := compareAuthsBcrypt(toHash, userid)
	if err != nil {
		log.Fatal(err)
	}

	if authenticated == true {
		fmt.Println("AUTHENTICATED")
	} else {
		fmt.Println("NO MATCH FOUND")
	}
}

// AuthenticateBcryptAPI - to do the authentication I suppose.
func AuthenticateBcryptAPI(userid string, auth string) bool {
	authenticated, err := compareAuthsBcrypt(auth, userid)
	if err != nil {
		log.Fatal(err)
	}

	if authenticated == true {
		fmt.Println("AUTHENTICATED")
		return true
	}
	fmt.Println("NO MATCH FOUND")
	return false

}

func compareAuthsBcrypt(toCompare string, userid string) (bool, error) {
	authSlice, err := getAuths(userid) // get all of the auths into a slice
	if err != nil {
		return false, err
	}
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
	return authenticated, nil
}
