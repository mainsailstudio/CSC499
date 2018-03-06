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

// GetLoginState - load the user properties via email and return what to do next to the API
func GetLoginState(w http.ResponseWriter, r *http.Request) {
	var user UserLoginState
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" {
		http.Error(w, "The email is empty", 400)
	}
	userExists, userID := checkUserExists(user.Email)
	fmt.Println("User exists bool is", userExists)
	if userExists == true {
		lockString, passExpire := getUserLocksAndPass(userID)
		if passExpire == true {
			newUserLogin := UserLoginState{Email: user.Email, LoginState: "3", Locks: lockString}
			json.NewEncoder(w).Encode(newUserLogin)
		} else if passExpire != true && lockString != "No locks received" {
			newUserLogin := UserLoginState{Email: user.Email, LoginState: "2", Locks: lockString}
			json.NewEncoder(w).Encode(newUserLogin)
		} else {
			newUserLogin := UserLoginState{Email: user.Email, LoginState: "1", Locks: ""}
			json.NewEncoder(w).Encode(newUserLogin)
		}
		//json.NewEncoder(w).Encode(user)

	} else {
		http.Error(w, "The email does not exist in our records", 400)
	}
}

// LoginUser - the function that actually logs in the user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user UserLoginCheck
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" {
		http.Error(w, "The email is empty", 400)
	}

	userExists, userID := checkUserExists(user.Email)

	if userExists == true {
		fmt.Println("User login state is", user.LoginState)
		if user.LoginState == "3" {
			// if the user is logging in via dynauth
			authCorrect := dynauthcore.AuthenticateBcryptAPI(userID, user.Secret)
			if authCorrect {
				fmt.Println("Correctly authenticated via auths")
				token := issueJWT(user.Email) // sending the user's email to be a part of the jwt claim
				userSuccess := UserLoginSuccess{ID: userID, Email: user.Email, LoginState: user.LoginState, Token: token}
				json.NewEncoder(w).Encode(userSuccess)
			} else {
				fmt.Println("NOT authenticated via auths")
			}
		}

		if user.LoginState == "1" {
			// if the user is loggin in via a temp pass
			passwordCorrect := dynauthcore.TempPassAuth(userID, user.Secret)
			if passwordCorrect {
				fmt.Println("Correctly authenticated via a temp pass")
				token := issueJWT(user.Email) // sending the user's email to be a part of the jwt claim
				userSuccess := UserLoginSuccess{ID: userID, Email: user.Email, LoginState: user.LoginState, Token: token}
				json.NewEncoder(w).Encode(userSuccess)
			} else {
				fmt.Println("NOT authenticated via a temp pass")
			}
		}
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
		fmt.Println("There is no date")
		return true // date is expired
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
