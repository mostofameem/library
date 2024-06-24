package db

import "github.com/jmoiron/sqlx"

//"github.com/jmoiron/sqlx"

var Db *sqlx.DB

func GetDB() *sqlx.DB {
	return Db
}
