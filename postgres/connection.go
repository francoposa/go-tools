package postgres

import (
	"database/sql"
	"fmt"
)

// ConnectionConfig defines Postgres database connection information
type ConnectionConfig struct {
	Host                  string
	Port                  uint16
	Username              string
	Password              string
	Database              string
	ApplicationName       string
	ConnectTimeoutSeconds int
	SSLMode               string
}

// BuildConnectionURI builds a connection string for lib/pq from ConnectionConfig.
// If a missing or invalid field is provided, an error is returned.
func BuildConnectionURI(cc ConnectionConfig) string {
	auth := ""
	if cc.Username != "" || cc.Password != "" {
		auth = fmt.Sprintf("%s:%s@", cc.Username, cc.Password)
	}
	url := fmt.Sprintf(
		"postgres://%s%s:%d/%s?application_name=%s&connect_timeout=%d&sslmode=%s",
		auth,
		cc.Host,
		cc.Port,
		cc.Database,
		cc.ConnectTimeoutSeconds,
		cc.SSLMode,
	)
	return url
}

// MustConnect mimics sqlx MustConnect, but for sql.DB, using ConnectionConfig
// Opens connection to a DB, pings, and panics on error
func MustConnect(cc ConnectionConfig) *sql.DB {
	uri := BuildConnectionURI(cc)

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

// MustOpen mimics sqlx MustOpen, but for sql.DB, using ConnectionConfig
// Opens connection to a DB and panics on error
func MustOpen(cc ConnectionConfig) *sql.DB {
	uri := BuildConnectionURI(cc)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	return db
}
