/*
	Title:	Storing package
	Author:	Connor Peters
	Date:	2/12/2018
	Desc:
*/

package dynauthcore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver helper
)

//	StorePerms - to store a slice of hashed permutations into a MySQL database.
//
func StorePerms(toInsert []string) {
	db, err := sql.Open("mysql", "root:root@/dynauth-dev")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	prepareStatement := "INSERT INTO hashes VALUES("
	for i := 1; i < len(toInsert); i++ {
		prepareStatement += "?), ("
	}
	prepareStatement += "?)"

	stmtIns, err := db.Prepare(prepareStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	dataPrepared := []interface{}{}
	for i := 0; i < len(toInsert); i++ {
		dataPrepared = append(dataPrepared, toInsert[i])
	}
	_, err = stmtIns.Exec(dataPrepared...) // adds all data in the slice as a separate argument BEAUTIFUL
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
