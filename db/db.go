package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const connectionTimeout = 10 * time.Second

const (
	host     = "localhost"
	port     = 5432
	user     = "koochooloo"
	password = "koochooloo"
	dbname   = "cinema"
)

// New creates a new postgres connection and tests it.
func New() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db, nil
}
