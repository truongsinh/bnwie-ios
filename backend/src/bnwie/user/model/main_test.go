package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func prepareDB() *sql.DB {
	if db != nil {
		return db
	}
	d, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	db = d
	return db
}
