package svc

import "github.com/jmoiron/sqlx"

//
// Provides authentication services
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

// NewAuthSvc builds and returns a new, properly instantiated AuthSvc
func NewAuthSvc(db *sqlx.DB, guids *Guids) AuthSvc {
	userDao := newUserDao(db, guids)
	return AuthSvc{userDao}
}

type AuthSvc struct {
	userDao *UserDao
}

// func (svc AuthSvc) SavePassword(db *sqlx.DB, email string, password string) (err error) {
// 	// Find the userId for the given email
// 	var userId int
// 	err = db.Get(&userId, findUserIdByEmail, email)
// 	svc.userData.findIDAndPasswordByEmail(email string)
//
// 	if err != nil {
// 		return err
// 	}
// 	if userId == 0 {
// 		err = errors.New("email not found")
// 	} else {
// 		var hashedPassword string
// 		hashedPassword, err = misc.HashPasswordDefaultCost(password)
// 		if err == nil {
// 			_, err = db.Exec(updatePasswordByUserId, hashedPassword, userId)
// 		}
// 	}
//
// 	return err // err may be nil
// }

// Authenticate method... When successful, the returned user will have a non-negative UserId.
func (svc AuthSvc) Authenticate(email string, testPassword string) (user User, err error) {
	// Find the hashed password for the given email
	var rows *sql.Rows
	rows, err = svc.userDao.findUserIdAndPasswordByEmail(email)
	if err != nil {
		return
	}
	for rows.Next() {
		var userId int
		var hashedPassword sql.NullString
		err = rows.Scan(&userId, &hashedPassword)
		if err != nil {
			log.Println("AuthSvc.Authenticate() Error in Scan:", err)
		} else if userId == 0 {
			log.Println("AuthSvc.Authenticate() user not found for email:", email)
		} else if hashedPassword.Valid {
			if misc.ComparePassword(testPassword, hashedPassword.String) {
				// We found a match so find the user
				user, err = svc.userDao.findById(userId)
			} else {
				log.Println("AuthSvc.Authentic() Password mismatched for email:", email)
			}
		}
		rows.Close()
	}

	// return the user and error
	return user, err
}
