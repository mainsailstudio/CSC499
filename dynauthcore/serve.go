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
	"fmt"
	"log"

	"github.com/shogo82148/go-shuffle"
)

//ServeLocks - to query the database and return the user's locks in a string
func ServeLocks(userid string) string {
	locks := getLocks(userid)
	fmt.Println("Initial locks are", locks)
	var lockString string
	shuffle.Slice(locks)
	fmt.Println("Shuffled locks are", locks)
	for i := range locks {
		lockString += locks[i]
	}
	return lockString
}

//ServeLockSlice - to query the database and return the user's locks in a slice
func ServeLockSlice(userid string) []string {
	locks := getLocks(userid)
	shuffle.Strings(locks)
	return locks
}

func getLocks(userid string) []string {
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
		log.Fatal(err)
	}
	defer locks.Close()
	for locks.Next() {
		var lockInfo string
		err := locks.Scan(&lockInfo)
		if err != nil {
			log.Fatal(err)
		}
		lockSlice = append(lockSlice, lockInfo)
	}
	err = locks.Err()
	if err != nil {
		log.Fatal(err)
	}

	return lockSlice
}
