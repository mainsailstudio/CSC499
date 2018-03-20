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

// insertConfigActivity - inserts the user's account configuration activity on the front-end
func insertConfigActivity(log ConfigActivity) error {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return errors.New("Unable to open the database connecting in the insertConfigActivity log")
	}
	defer db.Close()

	fmt.Println("User id is", log.UserID)
	fmt.Println("TotalCreation time is", log.TotalTime)
	fmt.Println("Avg length is", log.AvgSecretLength)

	initUser := "INSERT INTO testConfigLog (id, userid, totalCreationTime, avgSecretLength) VALUES (DEFAULT, ?, ?, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		return errors.New("Issue preparing to log user's config activity")
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(log.UserID, log.TotalTime, log.AvgSecretLength)
	if err != nil {
		errMessage := fmt.Sprintf("There was an issue executing the query to insert a new config activity log.\nError is", err)
		return errors.New(errMessage)
	}

	return nil
}

// insertLoginActivity - inserts the user's activity on the front-end
func insertLoginActivity(log LoginActivity) error {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return errors.New("Unable to open the database connecting in the insertLogActivity log")
	}
	defer db.Close()

	initUser := "INSERT INTO testLoginLog (id, userid, loginTime, failures, refreshes, secretLength) VALUES (DEFAULT, ?, ?, ?, ?, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		return errors.New("Issue preparing to log user's login activity")
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(log.UserID, log.LoginTime, log.Failures, log.Refreshes, log.SecretLength)
	if err != nil {
		errMessage := fmt.Sprintf("There was an issue executing the query to insert a new login activity log.\nError is", err)
		return errors.New(errMessage)
	}

	return nil
}
