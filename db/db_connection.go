package db

import (
	"database/sql"
	"fmt"
	"log"

	//mySQL drive
	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/video")
	if err == nil {
		fmt.Println("Database is connected")
	} else {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Println(":)")
	return db
}
