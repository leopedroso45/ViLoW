package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"log"
)

/*User object */
type User struct {
	IDUser, NameUser, PasswordUser, AgeUser string
}

/*GetUser User object */
func GetUser(name, password string) (User, error) {
	var con *sql.DB
	con = CreateCon()
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
	con = CreateCon()

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
	con = CreateCon()

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
