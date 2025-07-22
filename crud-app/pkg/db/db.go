package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

func DBConnection(host, port, user, password, dbname string) error {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
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
	return nil
}

func GetDBConnection() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}

func CloseDBConnection() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
