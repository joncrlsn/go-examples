package svc

//
// Provides private methods for querying and changing the user table
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

func newUserDao(db *sqlx.DB, guids *Guids) *UserDao {
	return &UserDao{db, guids}
}

// UserDao is a data access object
type UserDao struct {
	Db    *sqlx.DB
	guids *Guids
}

// User is a data object that happens to represents a row in the user table
type User struct {
	UserID    int64  `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

const (
	userInsertSQL                   = `INSERT INTO user (user_id, first_name, last_name, email) VALUES (:UserID, :FirstName, :LastName, :Email);`
	userUpdateSQL                   = `UPDATE user SET first_name = :FirstName, last_name = :LastName, email = :Email) WHERE user_id = :UserID;`
	userFindByEmailSQL              = `SELECT user_id, first_name, last_name, email FROM user WHERE email = $1;`
	userFindByIDSQL                 = `SELECT user_id, first_name, last_name, email FROM user WHERE user_id = $1;`
	userFindIDAndPasswordByEmailSQL = "SELECT user_id, password FROM user WHERE email = $1"
	userIDFindIDByEmailSQL          = "SELECT user_id FROM user WHERE email = $1"
	userUpdatePasswordByUserIDSQL   = "UPDATE user SET password = $1 WHERE user_id = $2"
)

func (dao *UserDao) create(user *User) error {
	if user.UserID != 0 {
		return errors.New("When creating a user, UserID must be zero")
	}
	user.UserID = dao.guids.next()
	_, err := dao.Db.NamedExec(userInsertSQL, user)
	return err
}

func (dao *UserDao) update(user User) error {
	if user.UserID == 0 {
		return errors.New("When updating a user, UserID cannot be zero")
	}
	_, err := dao.Db.NamedExec(userUpdateSQL, user)
	return err
}

// FindByEmail returns a slice of User instances
func (dao *UserDao) findByEmail(email string) ([]User, error) {
	users := []User{}
	err := dao.Db.Select(&users, userFindByEmailSQL, email)
	return users, err
}

// FindById returns an instance for the given id.  UserID will be 0 if none found.
func (dao *UserDao) findByID(id int) (User, error) {
	var user User
	err := dao.Db.Get(&user, userFindByIDSQL, id)
	return user, err
}

func (dao *UserDao) findIDAndPasswordByEmail(email string) (*sql.Rows, error) {
	return dao.Db.Query(userFindIDAndPasswordByEmailSQL, email)
}
