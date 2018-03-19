/*
	Title:	Authentication package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	This is the package that takes in
*/

package dynauthcore

import (
	"database/sql"
	dbinfo "dbinfo"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

//AuthenticateScrypt - to do the authentication I suppose.
func AuthenticateScrypt(locks string, otp string, userid string, iterations int) {
	// first prep auth for comparison
	//salts := getSalts(userid)
	auths, err := getAuths(userid)
	if err != nil {
		log.Fatal(err)
	}
	authenticated := false
	for i := range auths {
		//toHash := locks + otp + salts[i]
		toHash := locks + otp + "salt"
		fmt.Println("========================")
		fmt.Println("Comparison number is", i+1)
		fmt.Println("To hash string is	", toHash)
		saltByte := []byte("salt")
		hashedOTP, err := hashScrypt(toHash, saltByte, iterations) // hash the prepped otp
		if err != nil {
			fmt.Println("Error encountered:", err)
		}

		fmt.Println("Hashed string is	", hashedOTP)
		fmt.Println("Auth to compare is	", auths[i])
		if hashedOTP == auths[i] {
			authenticated = true
			fmt.Println("AUTHENTICATED")
			break
		} else {
			authenticated = false
			fmt.Println("NO MATCH FOUND")
		}
	}
	if authenticated == true {
		fmt.Println("You were authenticated")
	} else {
		fmt.Println("You were NOT authenticated")
	}

	// fmt.Println("\n==================\nTEST\n==================")
	// str := "test"
	// salt := make([]byte, saltLength)
	// _, err := io.ReadFull(rand.Reader, salt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Salt is", salt)
	// bytes := []byte(str)
	// fmt.Println("Bytes is", bytes)
	// bytes += salt
	// fmt.Println("Both is",bytes)

	// // Converts string to sha2
	// h := sha256.New()                       // new sha256 object
	// h.Write(bytes)                          // data is now converted to hex
	// code := h.Sum(nil)                      // code is now the hex sum
	// codestr := hex.EncodeToString(code)     // converts hex to string
	// fmt.Println("Code string is", codestr)

	// salt := make([]byte, saltLength)
	// _, err := io.ReadFull(rand.Reader, salt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Salt byte is", salt)
	// saltString := fmt.Sprintf("%x", salt)
	// fmt.Println("Salt string is", saltString)
	// fmt.Println("Converted salt string is", []byte(saltString))
	// saltString2 := fmt.Sprintf("%x", []byte(saltString))
	// fmt.Println("Salt string is", saltString2)
	// fmt.Println()
	// testHash, err := scrypt.Key([]byte("test"), salt, iterations, 8, 1, 64)
	// fmt.Println("Hash byte is", testHash)
	// fmt.Printf("Hash string is %x ", testHash)
	// fmt.Println()

	// fmt.Println("\n==================\nCOMPARE\n==================")
	// compareTest, err := scrypt.Key([]byte("test"), []byte(saltString), iterations, 8, 1, 64)
	// fmt.Println("Hash byte is", compareTest)
	// fmt.Printf("Hash string is %x ", compareTest)
	// fmt.Println()
}

func hashScrypt(otp string, salt []byte, iterations int) (string, error) {
	otpByte := []byte(otp)
	hashedPasswordScrypt, err := scrypt.Key(otpByte, salt, iterations, 8, 1, 64)
	if err != nil {
		return "", errors.New("there was an issue hashing the string using scrypt")
	}
	hashedToString := fmt.Sprintf("%x", hashedPasswordScrypt)
	return hashedToString, nil
}

func compareAuthsString(toCompare string, userid string) bool {
	authSlice, err := getAuths(userid) // get all of the auths into a slice
	if err != nil {
		log.Fatal(err)
	}
	var authenticated bool
	fmt.Println("Auth slice is", authSlice)
	for i := range authSlice {
		fmt.Println("Compare number", i)
		authenticated = false
		if toCompare == authSlice[i] {
			authenticated = true
		}
	}
	return authenticated
}

func getAuths(userid string) ([]string, error) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return nil, errors.New("Opening the database connection for getAuths went wrong")
	}

	defer db.Close()
	authSlice := []string{}
	query := "SELECT auth FROM auth" + userid
	auths, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer auths.Close()
	for auths.Next() {
		var auth string
		err := auths.Scan(&auth)
		if err != nil {
			log.Fatal(err)
		}
		authSlice = append(authSlice, auth)
	}
	err = auths.Err()
	if err != nil {
		log.Fatal(err)
	}
	return authSlice, nil
}

func getSalts(userid string) ([]string, error) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return nil, errors.New("Opening the database connection for getSalts went wrong")
	}

	defer db.Close()
	lockSlice := []string{}
	query := "SELECT salt FROM auth" + userid
	locks, err := db.Query(query)
	if err != nil {
		return nil, errors.New("Getting the user's salts caused an error")
	}

	defer locks.Close()
	for locks.Next() {
		var lock string
		err := locks.Scan(&lock)
		if err != nil {
			log.Fatal(err)
		}
		lockSlice = append(lockSlice, lock)
	}
	err = locks.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lockSlice, nil
}

// this gets the user's temp locks that were stored when served to enable the authentication
func getTestTempLocks(userid string) (string, error) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return "", errors.New("Opening the database connection for getTestTempLocks went wrong")
	}
	defer db.Close()

	// select the temp pass of the user
	var lockString string
	err = db.QueryRow("SELECT locks FROM tempTestLocks WHERE userid = ?", userid).Scan(&lockString)
	if err != nil {
		log.Fatal(err)
	}
	return lockString, nil
}
