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
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser - create and insert a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user UserRegisterStart
	_ = json.NewDecoder(r.Body).Decode(&user)
	userExists, userID, err := checkUserExists(user.Email)
	if err != nil {
		http.Error(w, "Error encountered when checking if the user exists", 500)
	}
	_ = userID // userID should be empty here

	if userExists == false {
		registerUser(user.Email, user.TempPass)
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "This email already exists, please use a different email", 400)
	}
}

func checkUserExists(email string) (bool, string, error) {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return false, "", errors.New("Opening the database connection for checkUserExists went wrong")
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
	return exists, userID, nil
}

// registerUse - only takes email and tempPass for simple registration
func registerUser(email string, tempPass string) error {
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		return errors.New("Opening the database connection for registerUser went wrong")
	}
	defer db.Close()

	// This is where each unique user is created
	initUser := "INSERT INTO users (id, email) VALUES (DEFAULT, ?)"
	stmtIns, err := db.Prepare(initUser)
	if err != nil {
		return errors.New("Preparing to insert a new user failed")
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(email)
	if err != nil {
		return errors.New("Executing the query to insert a new user failed")
	}

	// select the userid of the user that was just created
	// nice example of a simple single row query
	var userid string
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&userid)
	if err != nil {
		return errors.New("Getting the last insert ID to finish inserting a new user has failed")
	}
	fmt.Println("User id selected was", userid)

	// This is where each unique user password is created
	initUserPass := "INSERT INTO tempPass (userid, pass, expireDate, init) VALUES (?, ?, ?, ?)"
	stmtInsPass, err := db.Prepare(initUserPass)
	if err != nil {
		return errors.New("Preparing to finish the insertion of a new user has failed")
	}
	defer stmtInsPass.Close()

	// create hashed password
	hashedPasswordBcrypt, err := bcrypt.GenerateFromPassword([]byte(tempPass), dynauthconst.BcryptIterations)
	if err != nil {
		return errors.New("Hashing the user's temporary password as a new user using Bcrypt has failed")
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
		return errors.New("Executing the query to finish the insertion of a new user has failed")

	}
	// confirmEmail(email)
	fmt.Println("Need to send the user an email here")
	return nil
}

// CreateUserContinue - to continue with the user creation, adding name and security level
// NOTE!
// Instead of inserting and selecting the data right in this function, it would be more ideal if it called a 'persistable' function that does that and returns an actual error so this can return an HTTP error. It is just more consistent that way
func CreateUserContinue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user UserRegisterCont
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("User is", user)
	dbinfo := dbinfo.Db()
	db, err := sql.Open(dbinfo[0], dbinfo[1]) // gets the database information from the dbinfo package and enters the returned slice values as arguments
	if err != nil {
		fmt.Println("Error encountered when connecting to the database in CreateUserContinue\nError is:", err)
	}
	defer db.Close()

	// This is where each unique user is created
	updateUser := "UPDATE users SET fname = ?, lname = ?, security = ? WHERE id = ?"
	updateUserPrep, err := db.Prepare(updateUser)
	if err != nil {
		http.Error(w, "Problem updating the user information", 500)
	}
	defer updateUserPrep.Close()

	_, err = updateUserPrep.Exec(user.Fname, user.Lname, user.SecurityLv, user.ID)
	if err != nil {
		http.Error(w, "Problem updating the user information", 500)
	}

	sqlString := "SELECT * FROM securitylevels WHERE id = " + user.SecurityLv
	securityJSON, err := GetJSONFromSQLSelect(sqlString)
	if err != nil {
		fmt.Println("Error selecting security schemes")
	}

	fmt.Fprintf(w, securityJSON) // prints to browser
}

/*
* This call sends an email to one recipient, using a validated sender address
* Do not forget to update the sender address used in the sample
 */
// func confirmEmail() {
// 	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
// 	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

// 	mj := mailjet.NewMailjetClient(publicKey, secretKey)

// 	param := &mailjet.InfoSendMail{
// 		FromEmail: "cpete4@u.brockport.edu",
// 		FromName:  "Bob Patrick",
// 		Recipients: []mailjet.Recipient{
// 			mailjet.Recipient{
// 				Email: "design@mainsailstudio.com",
// 			},
// 		},
// 		Subject:  "Hello World!",
// 		TextPart: "Hi there !",
// 	}
// 	res, err := mj.SendMail(param)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Success")
// 		fmt.Println(res)
// 	}
// }

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
