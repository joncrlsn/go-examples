package svc

//
// Provides private methods for querying and changing the user table
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

import "github.com/jmoiron/sqlx"

// NewAuthSvc builds and returns a new, properly instantiated AuthSvc
func NewUserSvc(db *sqlx.DB, guids *Guids) UserSvc {
	userDao := newUserDao(db, guids)
	return UserSvc{userDao}
}

type UserSvc struct {
	userDao *UserDao
}

// FindByEmail returns a slice of User instances
func (svc UserSvc) FindByEmail(email string) ([]User, error) {
	return svc.userDao.findByEmail(email)
}

func (svc UserSvc) Create(first, last, email string) error {
	user := User{}
	user.FirstName = first
	user.LastName = last
	user.Email = email
	return svc.userDao.create(&user)
}
