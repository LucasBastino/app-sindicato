package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func CreateConnection() {
	host := "ndk.h.filess.io"
	user := "sindicatoDB_settingcry"
	password := "databasefilessio"
	port := "3307"
	dbName := "sindicatoDB_settingcry"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("error trying to connect with database:", dbName)
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection pinging")
		log.Fatal(err.Error())
	}

	fmt.Println("Succesfully connected to database:", dbName)
	DB = db

}
