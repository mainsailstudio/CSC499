/*
	Title:	Storing package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:
*/

package dynauthcore

import (
	"database/sql"
	dbinfo "dbinfo"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver helper
)

// StoreAuthsWithSalts - to store a slice of hashed permutations into a MySQL database.
func StoreAuthsWithSalts(authsWithSalts [][][]byte, userid string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user auth table is created
	createTable := "CREATE TABLE auth" + userid + " (auth char(128) binary, salt char(64) binary)"
	_, err = db.Exec(createTable) // like lean cuisine no preperation needed

	// Prepare statement for inserting the user's auth into the new table
	prepareStatement := "INSERT INTO auth" + userid + " VALUES("
	// for loop adds all perms into prepared statement
	for i := 1; i < len(authsWithSalts); i++ {
		prepareStatement += "?, ?), ("
	}
	prepareStatement += "?, ?)"

	stmtIns, err := db.Prepare(prepareStatement)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	// casts the data to insert into a slice interface for variadic function inclusion below, quite elegant
	dataPrepared := []interface{}{}
	for i := 0; i < len(authsWithSalts); i++ {
		tempAuth := fmt.Sprintf("%x", authsWithSalts[i][0])
		fmt.Println("Temp auth is", tempAuth)
		tempSalt := fmt.Sprintf("%x", authsWithSalts[i][1])
		fmt.Println("Temp salt is", tempSalt)
		dataPrepared = append(dataPrepared, tempAuth, tempSalt)
	}
	_, err = stmtIns.Exec(dataPrepared...) // adds all data in the slice as a separate argument (variadic) BEAUTIFUL
	if err != nil {
		panic(err.Error())
	}
}

// StoreAuthsPlain - to store a slice of hashed permutations into a MySQL database.
func StoreAuthsPlain(auths []string, userid string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user auth table is created
	createTable := "CREATE TABLE auth" + userid + " (auth char(64) binary)"
	_, err = db.Exec(createTable) // like lean cuisine no preperation needed

	// Prepare statement for inserting the user's auth into the new table
	prepareStatement := "INSERT INTO auth" + userid + " VALUES("
	// for loop adds all perms into prepared statement
	for i := 1; i < len(auths); i++ {
		prepareStatement += "?), ("
	}
	prepareStatement += "?)"

	stmtIns, err := db.Prepare(prepareStatement)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	// casts the data to insert into a slice interface for variadic function inclusion below, quite elegant
	dataPrepared := []interface{}{}
	for i := 0; i < len(auths); i++ {
		dataPrepared = append(dataPrepared, auths[i])
	}
	_, err = stmtIns.Exec(dataPrepared...) // adds all data in the slice as a separate argument (variadic) BEAUTIFUL
	if err != nil {
		panic(err.Error())
	}
}

// StoreLocks function stores the user' locks
// Needs the user's locks in a slice of strings, the user's id as a string, and the lockType as a string
func StoreLocks(locks []string, userid string, lockType string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting the user's auth into the new table
	prepareStatement := "INSERT INTO locks VALUES("
	// for loop adds all perms into prepared statement
	for i := 1; i < len(locks); i++ {
		prepareStatement += "DEFAULT, ?, ?, ?), ("
	}
	prepareStatement += "DEFAULT, ?, ?, ?)"

	stmtIns, err := db.Prepare(prepareStatement)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	// casts the data to insert into a slice interface for variadic function inclusion below, quite elegant
	dataPrepared := []interface{}{}
	for i := 0; i < len(locks); i++ {
		dataPrepared = append(dataPrepared, userid)
		dataPrepared = append(dataPrepared, locks[i])
		dataPrepared = append(dataPrepared, lockType)
	}
	_, err = stmtIns.Exec(dataPrepared...) // adds all data in the slice as a separate argument (variadic) BEAUTIFUL
	if err != nil {
		panic(err.Error())
	}
}

// StoreUserInfo takes in userinfo and stores it
func StoreUserInfo(userid string, fname string, lname string, email string, phone string, securityLv string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user is created
	prepareStatement := "INSERT INTO users (id, fname, lname, email, phone, security) VALUES (?, ?, ?, ?, ?, ?)"
	// _, err = db.Exec(createTable) // like lean cuisine no preperation needed
	stmtIns, err := db.Prepare(prepareStatement)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	// userData := []interface{}{}
	// userData = append(userData, userid)
	// userData = append(userData, fname)
	// userData = append(userData, lname)
	// userData = append(userData, email)
	// userData = append(userData, phone)
	// userData = append(userData, securityLv)
	_, err = stmtIns.Exec(userid, fname, lname, email, phone, securityLv)
	if err != nil {
		panic(err.Error())
	}
}
