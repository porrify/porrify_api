package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(user string, password string, host string, port string, dbName string) *sql.DB {
	log.Println("initialize DB")
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) //TODO proper error handling instead of panic in your app
	}

	log.Println("DB conn", db)
	return db
}
