package model

import (
	"database/sql"

	"github.com/Smarp/socnet"
	_ "github.com/lib/pq"
)

func PrepareAll(db *sql.DB) {
	prepareUserUpsert(db)
	prepareUpdateSerivceUpsert(db)
}

type User struct {
	*socnet.Profile
	*ExtendedData
}

type ExtendedData struct {
	UserId      uint
	Description string   `binding:"required"`
	Provider    []string `binding:"required"`
	Consumer    []string `binding:"required"`
}
