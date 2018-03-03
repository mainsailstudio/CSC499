/*
	Title:	Server package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:	This is the package that takes in
*/

package dynauthcore

import (
	"database/sql"
	dbinfo "dbinfo"
	"errors"
	"math/rand"
	"time"
	//"github.com/shogo82148/go-shuffle"
)

//ServeLocks - to query the database and return the user's locks in a string
func ServeLocks(userid string, lockNum int) string {
	locks, err := GetLocks(userid)
	if err != nil {
		return ""
	}
	// shuffle.Slice(locks) - from imported package
	shuffle(locks) // from internet code
	locks = locks[:lockNum]

	var lockString string

	for i := range locks {
		lockString += locks[i]
	}
	return lockString
}

//ServeLockSlice - to query the database and return the user's locks in a slice
func ServeLockSlice(userid string, lockNum int) []string {
	locks, err := GetLocks(userid)
	if err != nil {
		return nil
	}
	if len(locks) < 1 {
		return nil
	}
	// shuffle.Slice(locks) - from imported package
	shuffle(locks) // from internet code

	locks = locks[:lockNum]

	return locks
}

func shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

// GetLocks queries the database and returns all of the user's locks into a slice
func GetLocks(userid string) ([]string, error) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	lockSlice := []string{}
	//selectQuery := "SELECT locks FROM locks WHERE userid = ?" + userid
	locks, err := db.Query("SELECT locksAre FROM locks WHERE userid = ?", userid)
	if err != nil {
		return nil, errors.New("No locks were receieved from the database, user must not have initialized them")
	}
	defer locks.Close()
	for locks.Next() {
		var lockInfo string
		err := locks.Scan(&lockInfo)
		if err != nil {
			return nil, errors.New("Locks weren't added to the slice properly for unknown reasons")
		}
		lockSlice = append(lockSlice, lockInfo)
	}
	// err = locks.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return lockSlice, nil
}
