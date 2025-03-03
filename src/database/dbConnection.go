package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DBInfo struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

var InfoDB = DBInfo{
	Host:     "ndk.h.filess.io",
	User:     "sindicatoDB_settingcry",
	Password: "databasefilessio",
	Port:     "3307",
	DBName:   "sindicatoDB_settingcry",
}

func CreateConnection() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", InfoDB.User, InfoDB.Password, InfoDB.Host, InfoDB.Port, InfoDB.DBName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("error trying to connect with database:", InfoDB.DBName)
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection pinging")
		log.Fatal(err.Error())
	}

	fmt.Println("Succesfully connected to database:", InfoDB.DBName)
	DB = db

}
