package user

import (
	"bnwie/user/model"
	"database/sql"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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

var ts *httptest.Server

func TestMain(m *testing.M) {
	fbToken = os.Getenv("BNWIE_FB_TOKEN")
	model.PrepareAll(prepareDB())

	r := gin.New()
	RegisterHandler(r.Group("user"))
	ts = httptest.NewServer(r)

	m.Run()
}
