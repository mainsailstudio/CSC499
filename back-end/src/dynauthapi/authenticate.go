/*
Title:	Registering the user via
	Author:	Connor Peters
	Date:	2/26/2018
	Desc:
*/

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// type Token struct {
// 	Token string `json:"token"`
// }

// GetAPIUsers - to get api users for testing
func GetAPIUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sqlString := "select * from users"
	usersDB, err := GetJSONFromSQL(sqlString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// json.NewEncoder(w).Encode(usersDB)
	fmt.Println(usersDB)
	fmt.Fprintf(w, usersDB) // prints to browser
}

// PostAPIAuth - to get api users for testing
func PostAPIAuth(w http.ResponseWriter, r *http.Request) {

	// going to assume they are authenticated here for testing and going to issue a valid JWT
	token, err := createJwtToken()
	if err != nil {
		log.Println("Error Creating JWT token", err)
		http.Error(w, "Something went horribly wrong", 400)
		return
	}

	jsonToken, _ := json.Marshal(token)
	fmt.Fprintf(w, string(jsonToken)) // prints to browser

}

// func mainJwt(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	_ = json.NewDecoder(r.Body).Decode(&user)

// 	token := user.(*jwt.Token)

// 	claims := token.Claims.(jwt.MapClaims)

// 	log.Println("Email: ", claims["email"])

// 	//return c.String(http.StatusOK, "you are on the top secret jwt page!")
// }

// func login(w http.ResponseWriter, r *http.Request) error {
// 	var user User
// 	_ = json.NewDecoder(r.Body).Decode(&user)

// 	email := user.Email
// 	password := user.TempPass

// 	// check username and password against DB after hashing the password
// 	if email == "test@test.com" && password == "password" {

// 		// create jwt token
// 		token, err := createJwtToken()
// 		if err != nil {
// 			log.Println("Error Creating JWT token", err)
// 			http.Error(w, "Something went horribly wrong", 400)
// 			return err
// 		}

// 		// return c.JSON(http.StatusOK, map[string]string{
// 		//     "message": "You were logged in!",
// 		//     "token": token,
// 		// })
// 		json.NewEncoder(w).Encode(token)

// 	}

// 	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
// }

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
