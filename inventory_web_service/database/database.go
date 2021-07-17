package database

import (
	"database/sql"
	"log"
	"time"
)

var DBConn *sql.DB

func SetupDatabases() {
	var err error
	DBConn, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/inventorydb")

	if err != nil {
		log.Fatal(err)
	}

	DBConn.SetMaxOpenConns(4)
	DBConn.SetMaxIdleConns(4)
	DBConn.SetConnMaxLifetime(60 * time.Second)
}
