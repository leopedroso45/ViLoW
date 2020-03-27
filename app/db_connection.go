package main

import (
	"database/sql"
	"fmt"
	"log"

	//mySQL drive
	_ "github.com/go-sql-driver/mysql"
)

/*CreateCon mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/videos")
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return db
}
