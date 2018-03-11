/*
	Title:	Authentication package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	Authenticates a user using the SHA3 hashing algorithm
*/

package dynauthcore

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	sha3 "golang.org/x/crypto/sha3"
)

// AuthenticateSHA3 - to perform command line SHA3 hash authentication for testing
func AuthenticateSHA3(locks string, otp string, userid string, iterations int) {
	// first prep auth for comparison
	toHash := locks + otp

	authenticated, err := compareAuthsSHA3(toHash, userid)
	if err != nil {
		log.Fatal(err)
	}

	if authenticated == true {
		fmt.Println("AUTHENTICATED")
	} else {
		fmt.Println("NO MATCH FOUND")
	}
}

// // AuthenticateSHA3API - to perform authentication for the restful API
// func AuthenticateSHA3API(userid string, auth string) bool {
// 	authenticated, err := compareAuthsBcrypt(auth, userid)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if authenticated == true {
// 		fmt.Println("AUTHENTICATED")
// 		return true
// 	}
// 	fmt.Println("NO MATCH FOUND")
// 	return false

// }

func compareAuthsSHA3(toCompare string, userid string) (bool, error) {
	// initialize a false authentication return
	authenticated := false

	authSlice, err := getAuths(userid) // get all of the auths into a slice
	if err != nil {
		return false, err
	}

	// ready the private key
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}

	for i := range authSlice {

		fmt.Println("Compare number", i)

		// A MAC with 32 bytes of output has 256-bit security strength -- if you use at least a 32-byte-long key.
		h := make([]byte, 32)
		d := sha3.NewShake256()
		// Write the key into the hash.
		d.Write(pkSecret)
		// Now write the data.
		d.Write([]byte(toCompare))
		d.Read(h)

		hashString := fmt.Sprintf("%x\n", h)

		// need to trim the hash since it contains a \n character that ruins comparison
		hashString = strings.ToLower(strings.Trim(hashString, " \r\n"))

		if strings.Compare(hashString, authSlice[i]) == 0 {
			authenticated = true
			break
		}

	}

	return authenticated, nil
}
