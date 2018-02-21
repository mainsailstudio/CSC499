/*
	Title:	Dynamic authentication testbed
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	Just used for cli testing of package dynauthcore
*/

package main

import (
	dynauthapi "dynauthapi"
	dynauthcore "dynauthcore"
	"fmt"
)

// program constants
const keyNum = 10        // this is the total amount of keys the user will want, be careful with this because the larger this number is, the factorially larger the amount of computations will be. Keep it > 30
const displayLockNum = 4 // this is the total amount of locks that will be displayed for dynamic authentication. Same goes for this, keep it small (> 7)
const hashIterations = 8 // currently using 10 for speed, but it is recommended to use at least 100,000 for server side
const hashLen = 64       // 64 equals 128 characters

func main() {
	fmt.Println("Starting API")
	dynauthapi.StartAPIPlease()
	fmt.Println("1. Add a new user (auths only)")
	fmt.Println("2. Authenticate a user")
	var selection int
	fmt.Scan(&selection)
	if selection == 1 {
		createUser()
	} else if selection == 2 {
		authenticateUser()
	} else {
		fmt.Println("Nada sir")
	}
} // end of main

func createUser() {
	// ask for num of lock/key combos (testing purposes)
	fmt.Println("The default number of lock/key combos as defined by the keyNum constant are", keyNum)
	fmt.Println("For testing purposes please enter in the number of lock/key combos you want: ")
	var numHash int
	fmt.Scan(&numHash)

	// get user data
	var userid string
	var fname string
	var lname string
	var email string
	var phone string
	var securityLv string
	fmt.Println("For testing purposes please enter in the userid (randomize please): ")
	fmt.Scan(&userid)
	fmt.Println("For testing purposes please enter in the firstname: ")
	fmt.Scan(&fname)
	fmt.Println("For testing purposes please enter in the lastname: ")
	fmt.Scan(&lname)
	fmt.Println("For testing purposes please enter in the email: ")
	fmt.Scan(&email)
	fmt.Println("For testing purposes please enter in the phone: ")
	fmt.Scan(&phone)
	fmt.Println("For testing purposes please enter in the security level: ")
	fmt.Scan(&securityLv)

	// store user data
	dynauthcore.StoreUserInfo(userid, fname, lname, email, phone, securityLv)
	fmt.Println("User info was stored")

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

	// store locks
	dynauthcore.StoreLocks(lockSlice, userid, "1") // using 1 as the locktype since it does nothing currently
	fmt.Println("User's locks were stored")

	// create and store auths
	lockPerms := dynauthcore.LimPerms(lockSlice, displayLockNum) // create the limited permutations for the locks from the dynauthcore permutations.go package
	keyPerms := dynauthcore.LimPerms(keySlice, displayLockNum)
	permsToHash := dynauthcore.CombinePerms(lockPerms, keyPerms) // create the perms to hash (should most likely be in a package eventually)
	fmt.Println("Perms to hash is", permsToHash)
	fmt.Println("Total number of permutations is", len(permsToHash))
	hashedPermsWithSalt := dynauthcore.HashPermsScrypt(permsToHash, hashIterations)
	//fmt.Println("Hashed perms is", hashedPermsWithSalt)
	//fmt.Println("Let's try to store them!")
	dynauthcore.StoreAuthsWithSalts(hashedPermsWithSalt, userid)
	fmt.Println("User's auths were stored")
	//hash(lockKeySlice)
}

func authenticateUser() {
	fmt.Println("Enter in user id")
	var userid string
	fmt.Scan(&userid)
	locks := dynauthcore.ServeLocks(userid) // receives the slice of locks from serve.go
	fmt.Println("Locks are", locks)
	var otp string
	fmt.Println("Enter in OTP")
	fmt.Scan(&otp)
	dynauthcore.AuthenticateScrypt(locks, otp, userid, hashIterations)

}
