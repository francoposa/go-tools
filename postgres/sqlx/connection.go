package sqlx

import (
	"github.com/jmoiron/sqlx"

	"github.com/francoposa/go-tools/postgres"
)

// MustOpen wraps sqlx MustOpen, using ConnectionConfig
// Opens connection to a DB and panics on error
func MustOpen(cc postgres.ConnectionConfig) *sqlx.DB {
	uri := postgres.BuildConnectionURI(cc)
	return sqlx.MustOpen("postgres", uri)
}

// MustConnect wraps sqlx MustConnect, using ConnectionConfig
// Opens connection to a DB, pings, and panics on error
func MustConnect(cc postgres.ConnectionConfig) *sqlx.DB {
	uri := postgres.BuildConnectionURI(cc)
	return sqlx.MustConnect("postgres", uri)
}
