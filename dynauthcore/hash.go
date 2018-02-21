/*
	Title:	Hashing package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	To finish the hashing process of the locks and keys
			First combine the lock and key permutation slices created by the "LimPerms" function in the permutation package
			Then call the appropriate hashing method (as of 2/12/2018 only bcrypt is supported)
*/

package dynauthcore

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

const saltLength = 32
const authLength = 64

//CombinePerms - to correctly concat 2 slices of permutations into 1.
//Needs 2 slices, 1 of locks and 1 of keys. It is assumed that they match up perfectly and will result in logic errors if they do not.
func CombinePerms(locks []string, keys []string) []string {
	combined := []string{} // assumes locks and keys are of same length (SHOULD ALWAYS BE)
	for i := 0; i < len(locks); i++ {
		combineString := locks[i] + keys[i]
		combined = append(combined, combineString)
	}
	return combined
}

//HashPermsBcrypt - takes in the slice to hash and the amount of iterations to use for bcrypt and returns a completely hashed slice of strings.
//Needs 1 slice to hash and the iteration number as an int.
func HashPermsBcrypt(toHash []string, iterations int) []string {
	hashed := []string{}
	for i := 0; i < len(toHash); i++ {
		hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(toHash[i]), iterations)
		if err != nil {
			panic(err)
		}
		hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
		hashed = append(hashed, hashedToString)
	}
	return hashed
}

//HashPermsScrypt - takes in the slice to hash and the amount of iterations to use for scrypt and returns a completely hashed slice of strings.
//Needs 1 slice to hash and the iteration number as an int.
func HashPermsScrypt(toHash []string, iterations int) [][][]byte {
	salts := [][]byte{}
	salts = createSalts(len(toHash))
	fmt.Println("Salt slice is", salts)
	authsWithSalts := [][][]byte{}
	for i := 0; i < len(toHash); i++ {
		//hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
		hashedPasswordScrypt, err := scrypt.Key([]byte(toHash[i]), salts[i], iterations, 8, 1, authLength)
		if err != nil {
			panic(err)
		}
		withSalts := [][]byte{}
		withSalts = append(withSalts, hashedPasswordScrypt)
		withSalts = append(withSalts, salts[i])
		authsWithSalts = append(authsWithSalts, withSalts)
	}
	return authsWithSalts
}

func createSalts(amountOfAuths int) [][]byte {
	saltSlice := [][]byte{}
	for i := 0; i < amountOfAuths; i++ {
		salt := make([]byte, saltLength)
		_, err := io.ReadFull(rand.Reader, salt)
		if err != nil {
			log.Fatal(err)
		}
		saltSlice = append(saltSlice, salt)
	}
	return saltSlice
}

// func HashPermSha3(toHash []string) []string {
// 	hashed := []string{}
// 	for i := 0; i < len(toHash); i++ {
// 		hashedPasswordSha3, err := sha3.New256(toHash[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 		hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
// 		hashed = append(hashed, hashedToString)
// 	}
// 	return hashed
// }
