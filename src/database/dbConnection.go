package database

import (
	"database/sql"
	"fmt"

	er "github.com/LucasBastino/app-sindicato/src/errors"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var AuthDB *sql.DB

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

// var InfoDB = DBInfo{
// 	Host:     "192.168.100.2",
// 	User:     "rootdos",
// 	Password: "panyquesoprueba",
// 	Port:     "3306",
// 	DBName:   "dbprueba",
// }

var InfoAuthDB = DBInfo{
	Host:     "jee2iw.h.filess.io",
	User:     "authsindicato_spentclear",
	Password: "2b2248e87783e410a977375914b1ab38f17e5db0",
	Port:     "3307",
	DBName:   "authsindicato_spentclear",
}

func CreateConnection() {
	// Host:     os.Getenv("HOST"),
	// User:     os.Getenv("USER"),
	// Password: os.Getenv("PASSWORD"),
	// Port:     os.Getenv("PORT"),
	// DBName:   os.Getenv("DB_NAME"),

	// alguien tiene acceso a esto? preguntar
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", InfoDB.User, InfoDB.Password, InfoDB.Host, InfoDB.Port, InfoDB.DBName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		er.DatabaseConnectionError.Msg = err.Error()
		// logearlo
		// hacer algo aca, capaz cambiar la forma en que se loggea, que primero entre al index, nose. Es para despues
		// panic(err)
		fmt.Println("error ingresando a la database")
	}

	err = db.Ping()
	if err != nil {
		er.DatabaseConnectionError.Msg = err.Error()
		// logearlo
		// hacer algo aca
		// panic(err)
		fmt.Println("error pingeando a la database")
	}

	authConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", InfoAuthDB.User, InfoAuthDB.Password, InfoAuthDB.Host, InfoAuthDB.Port, InfoAuthDB.DBName)
	authDB, err := sql.Open("mysql", authConnString)
	if err != nil {
		er.DatabaseConnectionError.Msg = err.Error()
		// logearlo
		// hacer algo aca, capaz cambiar la forma en que se loggea, que primero entre al index, nose. Es para despues
		// panic(err)
		fmt.Println("error ingresando a la auth database")
	}

	err = authDB.Ping()
	if err != nil {
		er.DatabaseConnectionError.Msg = err.Error()
		// logearlo
		// hacer algo aca
		// panic(err)
		fmt.Println("error pingeando a la auth database")
	}

	DB = db
	AuthDB = authDB
}
