package database_sql

import (
	"database/sql"

	"github.com/francoposa/go-tools/postgres"
	// Makes postgres driver available to Golang's database/sql package
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	_ "github.com/lib/pq"
)

// MustOpen mimics sqlx MustOpen, but for sql.DB, using ConnectionConfig
// Opens connection to a DB and panics on error
func MustOpen(cc postgres.ConnectionConfig) *sql.DB {
	uri := postgres.BuildConnectionURI(cc)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	return db
}

// MustConnect mimics sqlx MustConnect, but for sql.DB, using ConnectionConfig
// Opens connection to a DB, pings, and panics on error
func MustConnect(cc postgres.ConnectionConfig) *sql.DB {
	uri := postgres.BuildConnectionURI(cc)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		panic(err)
	}
	return db
}
