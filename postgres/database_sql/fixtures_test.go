package database_sql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/francoposa/go-tools/postgres"
)

func TestSetUpAndTearDownDB(t *testing.T) {
	superUserPGConfig := postgres.Config{
		Host:                  "localhost",
		Port:                  5432,
		Username:              "postgres",
		Password:              "",
		Database:              "postgres",
		ApplicationName:       "",
		ConnectTimeoutSeconds: 5,
		SSLMode:               "disable",
	}

	testDBName := RandomDBName("sql_tools_test")

	testDB, err := CreateDB(t, testDBName, superUserPGConfig)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that we created a DB with the correct name
	var connectedDBName string
	result := testDB.QueryRow(`SELECT current_catalog;`)
	err = result.Scan(&connectedDBName)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testDBName, connectedDBName)

	// Assert that attempting to create the same database fails
	duplicateTestDB, err := CreateDB(t, testDBName, superUserPGConfig)
	assert.Nil(t, duplicateTestDB)
	assert.NotNil(t, err)

	err = TearDownDB(t, testDB, superUserPGConfig)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the database was dropped
	superUserDB := MustConnect(superUserPGConfig)
	var output string
	result = superUserDB.QueryRow(
		fmt.Sprintf(
			`SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s'`,
			testDBName,
		),
	)
	err = result.Scan(&output)
	assert.EqualError(t, err, "sql: no rows in result set")

}
