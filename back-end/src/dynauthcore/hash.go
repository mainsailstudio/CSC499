/*
	Title:	Hashing package
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:	To finish the hashing process of the locks and keys
			First combine the lock and key permutation slices created by the "LimPerms" function in the permutation package
			Then call the appropriate hashing method (as of 2/12/2018 only bcrypt is supported)
*/

package dynauthcore

import (
	"bytes"
	"dynauthconst"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// CombinePerms - to correctly concat 2 slices of permutations into 1.
// Needs 2 slices, 1 of locks and 1 of keys. It is assumed that they match up perfectly and will result in logic errors if they do not.
func CombinePerms(locks []string, keys []string) []string {
	combined := []string{} // assumes locks and keys are of same length (SHOULD ALWAYS BE)
	for i := 0; i < len(locks); i++ {
		combineString := locks[i] + keys[i]
		combined = append(combined, combineString)
	}
	return combined
}

// HashPermsBcrypt - takes in the slice to hash and the amount of iterations to use for bcrypt and returns a completely hashed slice of strings.
// Needs 1 slice to hash and the iteration number as an int.
func HashPermsBcrypt(toHash []string) []string {
	hashed := []string{}
	for i := 0; i < len(toHash); i++ {
		hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(toHash[i]), dynauthconst.BcryptIterations)
		if err != nil {
			panic(err)
		}
		hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
		fmt.Println("Hashed string is", hashedToString)
		hashed = append(hashed, hashedToString)
	}
	return hashed
}
