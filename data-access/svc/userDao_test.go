package svc

import (
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	_ "github.com/mattn/go-sqlite3"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stvp/assert"
)

const (
	dbName = "__deleteme.db"
)

var (
	userDao *UserDao
	db      *sqlx.DB
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

	guids := NewGuids()
	userDao = newUserDao(db, guids)
}

func Test_UserFindByEmail(t *testing.T) {
	initialize(t)

	users, err := userDao.findByEmail("joncrlsn@gmail.com")
	assert.Nil(t, err, "error finding by email")
	assert.Equal(t, 1, len(users), "Expected 1 user returned")
	assert.Equal(t, "joncrlsn@gmail.com", users[0].Email, "Wrong email returned from db")
}
