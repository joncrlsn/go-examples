package svc

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
	"github.com/stvp/assert"
)

const (
	dbName = "__deleteme.db"
)

var (
	userData *UserDao
	db       *sqlx.DB
)

func initialize(t *testing.T) {
	// Synchronously setup the DB or migrate it to the latest

	allErrors, ok := migrate.ResetSync("sqlite3://"+dbName, "../migrations")
	if !ok {
		t.Fatal("Error resetting test db", allErrors)
	}

	var err error
	db, err = sqlx.Connect("sqlite3", dbName)
	if err != nil {
		t.Fatal("Error connectin to db", err)
	}

	userData = newUserDao(db)
}

func Test_UserFindByEmail(t *testing.T) {
	initialize(t)

	users, err := userData.findByEmail("joncrlsn@gmail.com")
	assert.Nil(t, err, "error finding by email")
	assert.Equal(t, 1, len(users), "Expected 1 user returned")
	assert.Equal(t, "joncrlsn@gmail.com", users[0].Email, "Wrong email returned from db")
}
