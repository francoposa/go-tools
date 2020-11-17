package sqlx

import (
	"testing"

	// Makes postgres driver available to Golang's database/sql package
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/francoposa/go-tools/postgres"
)

func TestMustOpen(t *testing.T) {
	config := postgres.ConnectionConfig{
		Host:                  "localhost",
		Port:                  5432,
		Username:              "postgres",
		Password:              "",
		Database:              "postgres",
		ApplicationName:       "",
		ConnectTimeoutSeconds: 0,
		SSLMode:               "disable",
	}
	db := MustOpen(config)
	assert.NotNil(t, db)
}

func TestMustConnect(t *testing.T) {
	config := postgres.ConnectionConfig{
		Host:                  "localhost",
		Port:                  5432,
		Username:              "postgres",
		Password:              "",
		Database:              "postgres",
		ApplicationName:       "",
		ConnectTimeoutSeconds: 0,
		SSLMode:               "disable",
	}
	db := MustConnect(config)
	assert.NotNil(t, db)
}
