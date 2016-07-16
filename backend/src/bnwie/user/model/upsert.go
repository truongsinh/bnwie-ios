package model

import (
	. "bnwie/log"
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var upsertUserStmt *sql.Stmt

func prepareUserUpsert(db *sql.DB) {
	query := `
	INSERT INTO "User" (
		"FacebookId",
		"FullName",
		"AvatarUrl",
		"Description"
	)
 	VALUES (
 		$1,
 		$2,
 		$3,
 		$4
	)
	ON CONFLICT ("FacebookId") DO UPDATE SET
		"FullName"  = EXCLUDED."FullName",
		"AvatarUrl" = EXCLUDED."AvatarUrl"
	RETURNING "Description", "Provider", "Consumer"
	`
	s, err := db.Prepare(query)
	if err != nil {
		Log.WithError(err).Fatal("Cannot prepare")
	}
	Log.Info("Prepared!")
	upsertUserStmt = s
}

var UpsertUserModel = func(u *User) error {
	i, err := strconv.Atoi(u.Profile.Id)
	if err != nil {
		return err
	}
	prov := ""
	cons := ""
	err = upsertUserStmt.
		QueryRow(i, u.FullName, u.AvatarUrl, u.Description).
		Scan(&u.Description, &prov, &cons)
	if err != nil {
		return err
	}
	if len(prov) > 2 {
		u.Provider = strings.Split(prov[1:len(prov)-1], ",")
	} else {
		u.Provider = []string{}
	}
	if len(cons) > 2 {
		u.Consumer = strings.Split(cons[1:len(cons)-1], ",")
	} else {
		u.Consumer = []string{}
	}
	return nil
}
