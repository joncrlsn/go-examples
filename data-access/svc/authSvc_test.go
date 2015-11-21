package svc

//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
)

// func Test_SavePassword(t *testing.T) {
// 	db := resetDb(t)
//
// 	// Save a new password
// 	err := SavePassword(db, "joncrlsn@gmail.com", "new-secret")
// 	assert.Nil(t, err, "error saving password for email %v", err)
//
// 	// See if we can authenticate using the password we just saved
// 	user, err2 := Authenticate(db, "joncrlsn@gmail.com", "new-secret")
// 	assert.Nil(t, err2, "error authenticating valid email address %v", err)
// 	assert.Equal(t, "joncrlsn@gmail.com", user.Email)
//
// 	// try an invalid email
// 	err = SavePassword(db, "xxxx@yyyy.com", "whatever")
// 	assert.NotNil(t, err, "Did not receive error when saving password for invalid email ")
// }
//
// func Test_Authenticate(t *testing.T) {
// 	db := resetDb(t)
//
// 	user, err := Authenticate(db, "xxxx@yyyy.com", "super-duper-secret")
// 	assert.Nil(t, err, "error authenticating bogus email address %v", err)
// 	assert.Equal(t, 0, user.UserId)
//
// 	user, err = Authenticate(db, "joncrlsn@gmail.com", "super-duper-secret")
// 	assert.Nil(t, err, "error authenticating valid email")
// 	assert.NotEqual(t, 0, user.UserId)
//
// 	user, err = Authenticate(db, "joncrlsn@gmail.com", "wrong-password")
// 	assert.Nil(t, err, "error authenticating valid email and wrong password")
// 	assert.Equal(t, 0, user.UserId)
// }

func resetDb(t *testing.T) (db *sqlx.DB) {
	// Synchronously setup the DB or migrate it to the latest
	allErrors, ok := migrate.ResetSync("sqlite3://"+dbName, "../migrations")
	if !ok {
		t.Fatal("Error resetting test db", allErrors)
	}

	db = sqlx.MustConnect("sqlite3", dbName)

	return db
}
