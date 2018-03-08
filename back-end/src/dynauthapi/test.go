/*
Title:	Registering the user via
	Author:	Connor Peters
	Date:	2/26/2018
	Desc:
*/

package api

import (
	"bytes"
	"database/sql"
	dbinfo "dbinfo"
	dynauthconst "dynauthconst"
	dynauthcore "dynauthcore"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type testUser struct {
	ID        string `json:"id"`
	Fname     string `json:"fname"`
	Lname     string `json:"Lname"`
	Email     string `json:"email"`
	Init      bool   `json:"init"`
	TestLevel int    `json:"testLevel"`
	Token     string `json:"token"`
}

func testTestAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This thing is working") // prints to browser
}

// issueJWT - to get api users for testing
func issueTestJWT(email string) string {

	pkSecret, err := ioutil.ReadFile("../../../private.ppk") // in form of byte
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(pkSecret))

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "Dynauth Test",
		"email": email,
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(pkSecret)
	if err != nil {
		fmt.Println("Error creating signed token")
		log.Fatal(err)
	}

	return tokenString
}

func testRegister(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user testUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("User email was", user.Email)
	userToken := issueTestJWT(user.Email)
	userExists, userID, userInit, userTestLevel := getTestUserInit(user.Email)

	if userExists {
		returnUser := testUser{ID: userID, Email: user.Email, Init: userInit, TestLevel: userTestLevel, Token: userToken}
		json.NewEncoder(w).Encode(returnUser)
	} else {
		http.Error(w, "This email is not pre-registered as a test user", 400)
	}
}

func testLogin(w http.ResponseWriter, r *http.Request) {

}

func getTestUserInit(email string) (bool, string, bool, int) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	exists := false
	// search to make sure this email doesn't already exist
	var userID string
	var init bool
	var testLevel int
	row := db.QueryRow("SELECT id, init, testLevel FROM testUsers where email = ?", email).Scan(&userID, &init, &testLevel)
	switch row {
	case sql.ErrNoRows:
		fmt.Println("No rows selected")
		exists = false
	default:
		exists = true
	}
	return exists, userID, init, testLevel
}

// registerUse - only takes email and tempPass for simple registration
func registerTestUser(email string, tempPass string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user is created
	initUser := "INSERT INTO users (id, email) VALUES (DEFAULT, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(email)
	if err != nil {
		panic(err.Error())
	}

	// select the userid of the user that was just created
	// nice example of a simple single row query
	var userid string
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&userid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User id selected was", userid)

	// This is where each unique user password is created
	initUserPass := "INSERT INTO tempPass (userid, pass, expireDate, init) VALUES (?, ?, ?, ?)"
	stmtInsPass, err := db.Prepare(initUserPass)
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsPass.Close()

	// create hashed password
	hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(tempPass), dynauthconst.BcryptIterations)
	if err != nil {
		panic(err)
	}
	tempPass = bytes.NewBuffer(hashedPasswordBcrypt).String()

	expireDate := time.Now().Local().AddDate(0, 0, 7)
	//timein := time.Now().Local().Add(time.Hour * time.Duration(Hours) +
	//time.Minute * time.Duration(Mins) +
	// time.Second * time.Duration(Sec))
	expireDate.Format("2006-01-02 15:04:05")
	fmt.Println("Expire date is", expireDate)
	fmt.Println("Hashed temp pass is", tempPass)
	_, err = stmtInsPass.Exec(userid, tempPass, expireDate, 0)
	if err != nil {
		panic(err.Error())
	}
	// confirmEmail()
	fmt.Println("Confirmation email was sent!!!")
}

// CreateUserContinue - to continue with the user creation, adding name and security level
func CreateTestUserContinue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user UserRegisterCont
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("User is", user)
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// This is where each unique user is created
	updateUser := "UPDATE users SET fname = ?, lname = ?, security = ? WHERE id = ?"
	updateUserPrep, err := db.Prepare(updateUser)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		http.Error(w, "Problem updating the user information", 500)
	}
	defer updateUserPrep.Close()

	_, err = updateUserPrep.Exec(user.Fname, user.Lname, user.SecurityLv, user.ID)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		http.Error(w, "Problem updating the user information", 500)
	}

	sqlString := "SELECT * FROM securitylevels WHERE id = " + user.SecurityLv
	securityJSON, err := GetJSONFromSQL(sqlString)
	if err != nil {
		fmt.Println("Error selecting security schemes")
	}

	fmt.Fprintf(w, securityJSON) // prints to browser
}

// GetLoginState - load the user properties via email and return what to do next to the API
func GetTestLoginState(w http.ResponseWriter, r *http.Request) {
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

func testGiveToken(w http.ResponseWriter, r *http.Request) {
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
func getTestUserLocksAndPass(userID string) (string, bool) {
	locks := dynauthcore.ServeLocks(userID, dynauthconst.DisplayLockNum)

	// check to see if locks came back
	if locks == "" {
		return locks, false
	}

	passExpire := checkUserTempPassDate(userID)
	// pass expire is either set to true or false
	return locks, passExpire
}
