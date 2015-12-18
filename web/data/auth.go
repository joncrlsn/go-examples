//
// Provides authentication methods
//
package data

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/joncrlsn/misc"
	"log"
)

const (
	findUserIdAndPasswordByEmail = "SELECT user_id, password FROM user WHERE email = $1"
	findUserIdByEmail            = "SELECT user_id FROM user WHERE email = $1"
	updatePasswordByUserId       = "UPDATE user SET password = $1 WHERE user_id = $2"
)

func SavePassword(db *sqlx.DB, email string, password string) (err error) {
	// Find the userId for the given email
	var userId int
	err = db.Get(&userId, findUserIdByEmail, email)
	if err != nil {
		return err
	}
	if userId == 0 {
		err = errors.New("email not found")
	} else {
		var hashedPassword string
		hashedPassword, err = misc.HashPassword(password)
		if err == nil {
			_, err = db.Exec(updatePasswordByUserId, hashedPassword, userId)
		}
	}

	return err // err may be nil
}

// When successful, the returned user will have a non-negative UserId.
func Authenticate(db *sqlx.DB, email string, testPassword string) (user User, err error) {
	// Find the hashed password for the given email
	var rows *sql.Rows
	rows, err = db.Query(findUserIdAndPasswordByEmail, email)
	if err != nil {
		return
	}
	for rows.Next() {
		var userId int
		var hashedPassword sql.NullString
		err = rows.Scan(&userId, &hashedPassword)
		if err != nil {
			log.Println("Error in Scan", err)
		} else if userId == 0 {
			//fmt.Println("userId == 0.  user not found")
		} else if hashedPassword.Valid {
			if misc.ComparePassword(testPassword, hashedPassword.String) {
				// We found a match so find the user
				user, err = UserFindById(db, userId)
			} else {
				log.Println("Password mismatched for email", email)
			}
		}
		rows.Close()
	}

	// return the user and error
	return
}
