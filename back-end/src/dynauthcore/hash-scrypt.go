/*
	Title:	Hashing package
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:	Hashing the OTPs with scrypt and a salt
*/

package dynauthcore

import (

	// "fmt"

	"fmt"

	"golang.org/x/crypto/scrypt"
)

// HashPermsScrypt - takes in the slice to hash and the amount of iterations to use for scrypt and returns a completely hashed slice of strings.
// Needs 1 slice to hash and the iteration number as an int.
func HashPermsScrypt(toHash []string) []string {
	hashed := []string{}
	for i := 0; i < len(toHash); i++ {
		hashedPasswordScrypt, err := scrypt.Key([]byte(toHash[i]), []byte("salt"), 16384, 8, 1, 32)
		if err != nil {
			fmt.Println("Error generating Scrypt with salt")
		}

		hashedToString := fmt.Sprintf("%x\n", hashedPasswordScrypt)

		fmt.Println("Hashed string is", hashedToString)
		hashed = append(hashed, hashedToString)
	}
	return hashed
}
