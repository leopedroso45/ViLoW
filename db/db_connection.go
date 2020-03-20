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
	db, err := sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/videos")
	if err == nil {
		fmt.Println("database is connected")

	} else {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	//defer db.Close()
	// make sure connection is available
	/* Detectar erro
	err = db.Ping()
	fmt.Println(err)
	*/
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	fmt.Println(":)")
	return db
}
