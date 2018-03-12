/*
	Title:	Hash SHA3
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:
*/

package dynauthcore

import (
	"fmt"
	"io/ioutil"
	"log"

	. "golang.org/x/crypto/sha3"
)

// HashPermsSHA3 - takes in the slice to hash and the amount of iterations to use for SHA3 and returns a completely hashed slice of strings.
// Needs 1 slice to hash and the iteration number as an int.
func HashPermsSHA3(toHash []string) []string {
	hashed := []string{}

	// get the private key from file
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}

	// iterates through toHash and hashes them all
	for i := 0; i < len(toHash); i++ {
		h := make([]byte, 32)
		d := NewShake256()
		// Write the key into the hash.
		d.Write(pkSecret)
		// Now write the data.
		d.Write([]byte(toHash[i]))
		// Read 32 bytes of output from the hash into h.
		d.Read(h)
		fmt.Printf("%x\n", h)

		hashString := fmt.Sprintf("%x\n", h)
		fmt.Println("Hash casted to string is", hashString)

		// add the new hash to the slice
		hashed = append(hashed, hashString)
	}
	return hashed
}

// HashPermsWithSaltSHA3 - takes in the slice to hash and the amount of iterations to use for SHA3 and returns a completely hashed slice of strings.
// Needs 1 slice to hash and the iteration number as an int.
func HashPermsWithSaltSHA3(toHash []string) [][]string {
	hashed := [][]string{}

	// get the private key from file
	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}

	// iterates through toHash and hashes them all
	for i := 0; i < len(toHash); i++ {
		h := make([]byte, 64)

		// generate salt
		s, err := getSaltString(32)
		if err != nil {
			log.Print(err)
		}

		// add salt to hash
		toHashWithSalt := toHash[i] + s
		fmt.Println("Random salt is", s)
		fmt.Println("To hash", toHashWithSalt)
		// create a new SHA3 256
		d := NewShake256()
		// Write the key into the hash.
		d.Write(pkSecret)
		// Now write the data.
		d.Write([]byte(toHashWithSalt))

		// Read 32 bytes of output from the hash into h.
		d.Read(h)
		fmt.Printf("Hashed is %x\n", h)

		hashString := fmt.Sprintf("%x\n", h)
		//fmt.Println("Hash casted to string is", hashString)

		// add the new hash to the slice
		keyPair := []string{}
		keyPair = append(keyPair, hashString)
		keyPair = append(keyPair, s)
		hashed = append(hashed, keyPair)
	}
	return hashed
}
