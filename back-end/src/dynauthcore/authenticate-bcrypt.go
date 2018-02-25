/*
	Title:	Authentication package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	This is the package that takes in
*/

package dynauthcore

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//AuthenticateBcrypt - to do the authentication I suppose.
func AuthenticateBcrypt(locks string, otp string, userid string, iterations int) {
	// first prep auth for comparison
	toHash := locks + otp
	fmt.Println("toHash is", toHash)
	// hashedOTP := hashBcrypt(toHash, iterations) // hash the prepped otp
	// fmt.Println("Hashed otp is", hashedOTP)
	if compareAuthsBcrypt(toHash, userid) == true {
		fmt.Println("AUTHENTICATED")
	} else {
		fmt.Println("NO MATCH FOUND")
	}
}

// func hashBcrypt(otp string, iterations int) string {
// 	hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(otp), iterations)
// 	if err != nil {
// 		panic(err)
// 	}
// 	hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
// 	return hashedToString
// }

func compareAuthsBcrypt(toCompare string, userid string) bool {
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
