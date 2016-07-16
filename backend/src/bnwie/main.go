package main

import (
	. "bnwie/log"
	"bnwie/user"
	"database/sql"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

func prepareDB() *sql.DB {
	d, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	return d
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), ginrus.Ginrus(Log, time.RFC3339, true))
	api := r.Group("api")
	user.RegisterHandler(api.Group("user"), prepareDB())

	r.Run("0.0.0.0:8080") // listen and server on 0.0.0.0:8080
}
