//
//
//
package data

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	//_ "github.com/mattn/go-sqlite3"
	//"github.com/joncrlsn/misc"
)

type User struct {
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

const (
	userInsertSql      = `INSERT INTO user (first_name, last_name, email) VALUES (:FirstName, :LastName, :Email);`
	userUpdateSql      = `UPDATE user SET first_name = :FirstName, last_name = :LastName, email = :Email) WHERE user_id = :UserId;`
	userFindByEmailSql = `SELECT user_id, first_name, last_name, email FROM user WHERE email = $1;`
	userFindByIdSql    = `SELECT user_id, first_name, last_name, email FROM user WHERE user_id = $1;`
)

// UserSave either inserts or updates the given User object to the database.
// If there is a userId, an update is done.  Otherwise, an insert.
func UserSave(db *sqlx.DB, user User) error {
	var err error
	if user.UserId == 0 {
		_, err = db.NamedExec(userInsertSql, user)
	} else if user.UserId > 0 {
		_, err = db.NamedExec(userUpdateSql, user)
	}
	return err
}

// UserFindByEmail returns an array User instances
func UserFindByEmail(db *sqlx.DB, email string) ([]User, error) {
	users := []User{}
	err := db.Select(&users, userFindByEmailSql, email)
	return users, err
}

// UserFindById returns one User for the given id.  UserId will be 0 if none found.
func UserFindById(db *sqlx.DB, userId int) (User, error) {
	var user User
	err := db.Get(&user, userFindByIdSql, userId)
	return user, err
}
