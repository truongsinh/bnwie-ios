package user

import (
	"bnwie/user/model"
	"database/sql"
	"os"
	"testing"
	"time"
)

var fbToken string

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

func TestMain(m *testing.M) {
	fbToken = os.Getenv("BNWIE_FB_TOKEN")
	model.PrepareAll(prepareDB())
	time.Sleep(1 * time.Second)
	println("PREPARED!!!")
	time.Sleep(1 * time.Second)
	m.Run()
}
