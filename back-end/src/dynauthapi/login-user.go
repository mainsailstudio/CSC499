/*
	Title:	Login the user
	Author:	Connor Peters
	Date:	2/26/2018
	Desc:
	Notes:	If the user has a tempPass set, they will be given an option to login with that
*/

package api

import (
	"database/sql"
	dbinfo "dbinfo"
	dynauthconst "dynauthconst"
	dynauthcore "dynauthcore"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	// . "github.com/mailjet/mailjet-apiv3-go"
)

// LoginUserFromEmail - load the user properties via email and return what to do next to the API
func LoginUserFromEmail(w http.ResponseWriter, r *http.Request) {
	var user UserLogin
	_ = json.NewDecoder(r.Body).Decode(&user)
	userExists, userID := checkUserExists(user.Email)
	if userExists == true {
		lockString, passExpire := getUserLocksAndPass(userID)
		if passExpire == true {
			newUserLogin := UserLogin{Email: user.Email, SecurityLv: "3", Locks: lockString}
			json.NewEncoder(w).Encode(newUserLogin)
		} else if passExpire != true && lockString != "" {
			newUserLogin := UserLogin{Email: user.Email, SecurityLv: "2", Locks: lockString}
			json.NewEncoder(w).Encode(newUserLogin)
		} else {
			newUserLogin := UserLogin{Email: user.Email, SecurityLv: "1", Locks: ""}
			json.NewEncoder(w).Encode(newUserLogin)
		}
		//json.NewEncoder(w).Encode(user)

	} else {
		http.Error(w, "The email does not exist in our records", 400)
	}
}

// gets the user's type of login
func getUserLocksAndPass(userID string) (string, bool) {
	locks := dynauthcore.ServeLocks(userID, dynauthconst.DisplayLockNum)

	// check to see if locks came back
	if locks == "" {
		return locks, false
	}

	passExpire := checkUserTempPassDate(userID)
	// pass expire is either set to true or false
	return locks, passExpire
}

// returns true or false if date is expired or not
func checkUserTempPassDate(userID string) bool {
	//fmt.Println("Did you even check?")
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var expireDateString string
	// search to make sure this email doesn't already exist
	err = db.QueryRow("SELECT expireDate FROM tempPass where userid = ?", userID).Scan(&expireDateString)
	if err != nil {
		return true
		// http.Error(w, "The email does not have a temppass set", 400)
	}

	timezone, _ := time.LoadLocation("EST")
	expireDate, err := time.ParseInLocation("2006-01-02 15:04:05", expireDateString, timezone)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	diff := expireDate.Sub(time.Now())

	if diff < 0 {
		fmt.Println("Date is expired")
		return true // date is expired
	}
	fmt.Println("Date is NOT expired!!")
	return false // date isn't expired
	// now := time.Now()

	// fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	// longTimeAgo := time.Date(2010, time.May, 18, 23, 0, 0, 0, time.UTC)

	// // compare time with time.Equal()

	// sameTime := longTimeAgo.Equal(now)

	// fmt.Println("longTimeAgo equals to now ? : ", sameTime)

	// // calculate the time different between today
	// // and long time ago

	// diff := now.Sub(longTimeAgo)

	// // convert diff to days
	// days := int(diff.Hours() / 24)

	// return expired
}
