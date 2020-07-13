package db

import (
	"database/sql"
	"log"

	/*Importing mysql driver */
	_ "github.com/go-sql-driver/mysql"
)

/*CreateCon open a mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/videos")
	if err != nil {
		log.Println("MySQL db is not connected")
		log.Println(err.Error())
		log.Fatal(err)
	}
	return db
}
