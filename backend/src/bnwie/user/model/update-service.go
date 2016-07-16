package model

import (
	. "bnwie/log"
	"database/sql"
	"errors"
	"strings"

	_ "github.com/lib/pq"
)

var updateSerivceStmt *sql.Stmt

func prepareUpdateSerivceUpsert(db *sql.DB) {
	query := `
	UPDATE "User" SET
		"Description"  = $2,
		"Provider" = $3,
		"Consumer" = $4
	WHERE "UserId" = $1
	`
	s, err := db.Prepare(query)
	if err != nil {
		Log.WithError(err).Fatal("Cannot prepare")
	}
	updateSerivceStmt = s
}

var UpdateServiceByIdModel = func(u *ExtendedData) error {
	prov := "{" + strings.Join(u.Provider, ",") + "}"
	cons := "{" + strings.Join(u.Consumer, ",") + "}"
	result, err := updateSerivceStmt.
		Exec(u.UserId, u.Description, prov, cons)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("User not found")
	}
	if aff != 1 {
		return errors.New("More than one user")
	}
	return nil
}
