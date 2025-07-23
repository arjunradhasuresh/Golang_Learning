package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DBConnection(host, port, user, password, dbname string) error {
	fmt.Println("hiiii")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("Database connection established successfully")
	return err
}

func GetDBConnection() *sql.DB {
	return db
}

func CloseDBConnection() {
	defer db.Close()
}
