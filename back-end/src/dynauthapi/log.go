/*
	Title:	Log user activity for the API
	Author:	Connor Peters
	Date:	3/18/2018
	Desc:	Logs the user activity depending on if they are using a password or dynauth
			This is used mostly for testing, and records the amount of failures, the amount of refreshes, and the total length of the string that they used to authenticate
*/

package api

import (
	"database/sql"
	dbinfo "dbinfo"
	"errors"
	"fmt"
)

// insertLogActivity - inserts the user's activity on the front-end
func insertLogActivity(log LogActivity) error {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return errors.New("Unable to open the database connecting in the insertLogActivity log")
	}
	defer db.Close()

	fmt.Println("Receieved user ID is", log.UserID)
	fmt.Println("Receieved LoginTime", log.LoginTime)
	fmt.Println("Receieved Failures", log.Failures)
	fmt.Println("Receieved Refreshes", log.Refreshes)
	fmt.Println("Receieved SecretLength", log.SecretLength)

	initUser := "INSERT INTO testLog (id, userid, loginTime, failures, refreshes, secretLength) VALUES (DEFAULT, ?, ?, ?, ?, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		return errors.New("Preparing to log user activity")
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(log.UserID, log.LoginTime, log.Failures, log.Refreshes, log.SecretLength)
	if err != nil {
		errMessage := fmt.Sprintf("There was an issue executing the query to insert a new activity log.\nError is", err)
		return errors.New(errMessage)
	}

	return nil
}
