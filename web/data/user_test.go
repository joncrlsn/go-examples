package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
	"github.com/stvp/assert"
	"testing"
)

const (
	dbName = "__deleteme.db"
)

func Test_UserFindByEmail(t *testing.T) {
	// Synchronously setup the DB or migrate it to the latest
	allErrors, ok := migrate.ResetSync("sqlite3://"+dbName, "../migrations")
	if !ok {
		t.Fatal("Error resetting test db", allErrors)
	}

	//
	db, err := sqlx.Connect("sqlite3", dbName)
	if err != nil {
		t.Fatal("Error connectin to db", err)
	}

	users, err := UserFindByEmail(db, "joncrlsn@gmail.com")
	assert.Nil(t, err, "error finding by email")
	assert.Equal(t, 1, len(users), "Expected 1 user returned")
	assert.Equal(t, "joncrlsn@gmail.com", users[0].Email, "Wrong email returned from db")
}
