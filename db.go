package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://localhost:5432/vaccination?sslmode=disable"
)

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}
	fmt.Println("ping: ", DB.Ping())
	// defer DB.Close()
}
