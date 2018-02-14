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

	_ "github.com/go-sql-driver/mysql" // mysql driver helper
)

//StorePerms - to store a slice of hashed permutations into a MySQL database.
func StorePerms(toInsert []string, userid string) {
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
	for i := 1; i < len(toInsert); i++ {
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
	for i := 0; i < len(toInsert); i++ {
		dataPrepared = append(dataPrepared, toInsert[i])
	}
	_, err = stmtIns.Exec(dataPrepared...) // adds all data in the slice as a separate argument (variadic) BEAUTIFUL
	if err != nil {
		panic(err.Error())
	}
}
