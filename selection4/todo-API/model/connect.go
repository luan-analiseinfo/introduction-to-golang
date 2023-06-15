package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var conn *sql.DB

func Connect() *sql.DB {
	sql.Drivers()
	db, err := sql.Open("mysql", "go_user:go_pass@tcp(localhost:3306)/go_database")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to the database")
	conn = db
	return db
}
