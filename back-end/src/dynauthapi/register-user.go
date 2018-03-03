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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// . "github.com/mailjet/mailjet-apiv3-go"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser - create and insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Fname != "" || user.Lname != "" || user.Phone != "" || user.LockNum != "" || user.KeyNum != "" || user.Security != nil {
		fmt.Println("Would update user to match here")
		//createFullUser(user.Fname, user.Lname, user.Phone, user.LockNum, user.KeyNum,)
	} else if user.Email != "" && user.TempPass != "" {
		userExists, userID := checkUserExists(user.Email)
		_ = userID // userID should be empty here
		if userExists == false {
			registerUser(user.Email, user.TempPass)
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, "This email already exists, please use a different email", 400)
		}
	} else {
		http.Error(w, "There was an error with the register user api call, the fields did not match any method", 400)
	}
}

func checkUserExists(email string) (bool, string) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	exists := false
	// search to make sure this email doesn't already exist
	var userID string
	row := db.QueryRow("SELECT id FROM users where email = ?", email).Scan(&userID)
	switch row {
	case sql.ErrNoRows:
		fmt.Println("No rows selected")
		exists = false
	default:
		exists = true
	}
	return exists, userID
}

// registerUse - only takes email and tempPass for simple registration
func registerUser(email string, tempPass string) {
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

func createFullUser() {

}

/*
* This call sends an email to one recipient, using a validated sender address
* Do not forget to update the sender address used in the sample
 */
// func confirmEmail() {
// 	type Payload struct {
// 		Messages []struct {
// 			From struct {
// 				Email string `json:"Email"`
// 				Name  string `json:"Name"`
// 			} `json:"From"`
// 			To []struct {
// 				Email string `json:"Email"`
// 				Name  string `json:"Name"`
// 			} `json:"To"`
// 			Subject  string `json:"Subject"`
// 			TextPart string `json:"TextPart"`
// 		} `json:"Messages"`
// 	}

// 	data := Payload{
// 		Messages[]{
// 			From{
// 				Email: "cpete4@u.brockport.edu",
// 				Name:  "Dynauth Test",
// 			},
// 			To{
// 				Email: "design@mainsailstudio.com",
// 				Name:  "Tester",
// 			},
// 			Subject:  "Test emaillls",
// 			TextPart: "TextPart I guess",
// 		},
// 	}

// 	payloadBytes, err := json.Marshal(data)
// 	if err != nil {
// 		// handle err
// 	}
// 	body := bytes.NewReader(payloadBytes)

// 	req, err := http.NewRequest("POST", "https://api.mailjet.com/v3.1/send", body)
// 	if err != nil {
// 		// handle err
// 	}
// 	req.SetBasicAuth("edc2ab073e461e2a00cb67bc1e714eab", "dc3dcad32f6fc03a925d2b35bd3f99b6")
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		// handle err
// 	}
// 	defer resp.Body.Close()
// }
