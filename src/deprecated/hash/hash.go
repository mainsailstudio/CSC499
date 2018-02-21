/*	Title:	Dynamic authentication
	Author:	Connor Peters
	Date:	2/3/2018
	Desc:
*/

package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	dynauthcore "dynauthcore"

	"github.com/dchest/pbkdf2"
	"golang.org/x/crypto/bcrypt"
)

// program constants
const keyNum = 10         // this is the total amount of keys the user will want, be careful with this because the larger this number is, the factorially larger the amount of computations will be. Keep it > 30
const displayLockNum = 4  // this is the total amount of locks that will be displayed for dynamic authentication. Same goes for this, keep it small (> 7)
const hashIterations = 10 // currently using 10 for speed, but it is recommended to use at least 100,000 for server side
const hashLen = 64        // 64 equals 128 characters

func main() {
	// ask for num of lock/key combos (testing purposes)
	fmt.Println("The default number of lock/key combos as defined by the keyNum constant are", keyNum)
	fmt.Println("For testing purposes please enter in the number of lock/key combos you want: ")
	var numHash int
	fmt.Scan(&numHash)

	// initialize the 2d lock-key combo slice
	lockSlice := make([]string, numHash)
	keySlice := make([]string, numHash)

	// for loop that asks for locks and keys (testing purposes)
	for i := 0; i < numHash; i++ {
		// intialize the slice of this particular iteration of lockKeySlice
		var lock string
		var key string

		// start by getting the lock and putting it into the slice
		fmt.Print("Enter in lock number ", i, ": ")
		fmt.Scan(&lock)
		fmt.Println("Lock is: " + lock) // print lock
		lockSlice[i] = lock

		// next get the key and put it into the slice
		fmt.Print("Enter in key correlating to lock number ", i, ": ")
		fmt.Scan(&key)
		fmt.Println("Key is: " + key) // print lock
		keySlice[i] = key
	}
	lockPerms := dynauthcore.LimPerms(lockSlice, displayLockNum) // create the limited permutations for the locks from the dynauthcore permutations.go package
	keyPerms := dynauthcore.LimPerms(keySlice, displayLockNum)
	permsToHash := combinePerms(lockPerms, keyPerms) // create the perms to hash (should most likely be in a package eventually)
	fmt.Println("Perms to hash is", permsToHash)
	fmt.Println("Total number of permutations is", len(permsToHash))
	//hashedPermsWithSalt := hashPermsWithSalt(permsToHash)
	//fmt.Println("Hashed perms is", hashedPermsWithSalt)
	//hash(lockKeySlice)

} // end of main

func combinePerms(locks []string, keys []string) []string {
	combined := []string{} // assumes locks and keys are of same length (SHOULD ALWAYS BE)
	for i := 0; i < len(locks); i++ {
		combineString := locks[i] + keys[i]
		combined = append(combined, combineString)
	}
	return combined
}

func hashPermsWithSalt(toHash []string) []string {
	//hashed := make([]string, len(locks)) // assumes locks and keys are of same length (SHOULD ALWAYS BE)
	hashed := []string{}
	for i := 0; i < len(toHash); i++ {
		// salt := make([]byte, 32)
		// if _, err := rand.Reader.Read(salt); err != nil {
		// 	panic("random reader failed")
		// }

		// to just use a set salt for testing
		// stringSalt := "randomsaltforrealz"
		// salt := []byte(stringSalt)
		//	hashedPasswordPBKDF2 := pbkdf2.WithHMAC(sha256.New, []byte(toHash[i]), salt, hashIterations, hashLen)
		hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(toHash[i]), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		hashedToString := bytes.NewBuffer(hashedPasswordBcrypt).String()
		//hashedToString := string(hashedPasswordPBKDF2[:])
		hashed = append(hashed, hashedToString)
	}
	return hashed
}

func hash(lks [][]string) {

	for i := 0; i < len(lks); i++ {
		var OTP string
		OTP = lks[i][1]
		fmt.Println("OTP is", OTP)

		// Hashing the password using bcrypt with the default cost of 10
		hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(OTP), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		// Hashing password using PBKDF2 and salt

		// to make an actually random salt
		salt := make([]byte, 32)
		if _, err := rand.Reader.Read(salt); err != nil {
			panic("random reader failed")
		}

		// to just use a set salt for testing
		// stringSalt := "randomsaltforrealz"
		// salt := []byte(stringSalt)
		hashedPasswordPBKDF2 := pbkdf2.WithHMAC(sha256.New, []byte(OTP), salt, hashIterations, hashLen)

		fmt.Println("Bcrypt hash is:", string(hashedPasswordBcrypt))
		fmt.Println("PBKDF2 hash is:")
		fmt.Printf("%x", hashedPasswordPBKDF2)
		fmt.Println()

	}

} // end of hash
