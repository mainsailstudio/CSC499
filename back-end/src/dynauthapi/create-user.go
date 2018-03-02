/*	Title:	RESTful API using Mux
	Author:	Connor Peters
	Date:	2/3/2018
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
	"os"
	"github.com/mailjet/mailjet-apiv3-go"
	
	"golang.org/x/crypto/bcrypt"
)

// CreateUser - create and insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Fname != "" || user.Lname != "" || user.Phone != "" || user.LockNum != "" || user.KeyNum != "" || user.Security != nil {
		fmt.Println("Would update user to match here")
		// updateFullUser
	} else if user.Fname == "" || user.Lname == "" || user.Phone == "" || user.LockNum == "" || user.KeyNum == "" || user.Security == nil && user.Email != "" && user.TempPass != "" {
		registerUser(user.Email, user.TempPass)
	} else {
		fmt.Println("There was an error with the register user api call, the fields did not match any method")
	}
	json.NewEncoder(w).Encode(user)
	fmt.Println("New user raw is", user, user.Security)
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
	initUserPass := "INSERT INTO tempPass (userid, pass) VALUES (?, ?)"
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

	fmt.Println("Hashed temp pass is", tempPass)
	_, err = stmtInsPass.Exec(userid, tempPass)
	if err != nil {
		panic(err.Error())
	}
	confirmEmail()
	fmt.Println("Confirmation email was sent!!!")
}

func insertUser() {
	// make sure to insert user without IDs

}


/*
* This call sends an email to one recipient, using a validated sender address
* Do not forget to update the sender address used in the sample
 */
func confirmEmail() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	param := &mailjet.InfoSendMail{
		FromEmail: "cpete4@u.brockport.edu",
		FromName:  "Bob Patrick",
		Recipients: []mailjet.Recipient{
			mailjet.Recipient{
				Email: "design@mainsailstudio.com",
			},
		},
		Subject:  "Hello World!",
		TextPart: "Hi there !",
	}
	res, err := mj.SendMail(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
