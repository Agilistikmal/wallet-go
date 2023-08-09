package database

import (
	"database/sql"
	"time"
)

func New() *sql.DB {
	db, err := sql.Open("mysql", "root:rootpw@tcp(localhost:3306)/wallet")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
