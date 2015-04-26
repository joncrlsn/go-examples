//
// Provides an example of the jmoiron/sqlx data mapping library with sqlite
//
package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);
`

type User struct {
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  sql.NullString
}

func main() {
	// this connects & tries a simple 'SELECT 1', panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO user (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO user (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", "John", "Doe", "johndoeDNE@gmail.net", "supersecret")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &User{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO user (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &User{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com"})
	tx.Commit()

	// Query the database, storing results in a []User (wrapped in []interface{})
	people := []User{}
	db.Select(&people, "SELECT * FROM user ORDER BY first_name ASC")
	jane, jason := people[0], people[1]

	fmt.Printf("Jane: %#v\nJason: %#v\n", jane, jason)
	// User{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// User{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// You can also get a single result, a la QueryRow
	jason1 := User{}
	err = db.Get(&jason1, "SELECT * FROM user WHERE first_name=$1", "Jason")
	fmt.Printf("Jason: %#v\n", jason1)
	// User{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	users := []User{}
	err = db.Select(&users, "SELECT * FROM user ORDER BY email ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	jane, jason, john := users[0], users[1], users[2]

	fmt.Printf("Jane: %#v\nJason: %#v\nJohn: %#v\n", jane, jason, john)

	// Loop through rows using only one struct
	user := User{}
	rows, err := db.Queryx("SELECT * FROM user WHERE first_name LIKE 'J%'")
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", user)
	}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = db.NamedExec(`INSERT INTO user (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	_, err = db.NamedExec(`UPDATE user SET first_name=:first, last_name=:last WHERE first_name = 'Bin'`,
		map[string]interface{}{
			"first": "Ben",
			"last":  "Smith",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, err = db.NamedQuery(`SELECT * FROM user WHERE first_name=:fn`, map[string]interface{}{"fn": "Ben"})
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Ben: %#v\n", user)
	}

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = db.NamedQuery(`SELECT * FROM user WHERE first_name=:first_name`, jason)
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Jason: %#v\n", user)
	}

	// fetch one column from the db
	rows2, _ := db.Query("SELECT first_name FROM user WHERE last_name = 'Smith'")
	// iterate over each row
	for rows2.Next() {
		var firstName string // if nullable, use the NullString type
		err = rows2.Scan(&firstName)
		fmt.Printf("Ben: %s\n", firstName)
	}
}
