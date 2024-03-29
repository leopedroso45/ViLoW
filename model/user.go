package model

import (
	"ViLoW/db"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
)

/*User object */
type User struct {
	IDUser, NameUser, PasswordUser, AgeUser string
}

/*UserG object */
type UserG struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}

/*GetUserFromID receive the userID from the video */
func GetUserFromID(idu int) (User, error) {
	var con *sql.DB
	con = db.CreateCon()
	var user User

	sqlst := fmt.Sprintf("SELECT user.id_user, user.name_user FROM user WHERE user.id_user = %v", idu)
	row, err := con.Query(sqlst)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&user.IDUser, &user.NameUser)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)

	}
	return user, err
}

/*GetUser User object */
func GetUser(name, password string) (User, error) {
	var con *sql.DB
	con = db.CreateCon()
	var user User
	encPassword := encrypting(password)
	sqlst := `SELECT user.id_user, user.name_user FROM user WHERE user.name_user= '` + name + `' AND user.password_user= '` + encPassword + `'`
	row := con.QueryRow(sqlst)
	err := row.Scan(&user.IDUser, &user.NameUser)
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned")
		return user, err
	case nil:
		return user, nil
	default:
		panic(err)
	}
}

/*CreateUser Receives an object of type User,
opens a connection to database and returns true
if no errors occur.*/
func CreateUser(user User) (result bool) {
	var con *sql.DB
	con = db.CreateCon()

	newPass := encrypting(user.PasswordUser)
	user.PasswordUser = newPass

	resultado, err := con.Query(`INSERT INTO user (name_user, age_user, password_user) VALUES ('` + user.NameUser + `', '` + user.AgeUser + `', '` + user.PasswordUser + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

/*clearDB Open a connection
to database and clears all
users already inserted.*/
func clearDBUser() (result bool) {
	var con *sql.DB
	con = db.CreateCon()

	resultado, err := con.Query("DELETE FROM user;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

func encrypting(senha string) string {
	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(senha))
	passwordCript := sha1Instance.Sum(nil)[:20]
	stringPasswordCript := hex.EncodeToString(passwordCript)
	return stringPasswordCript
}
