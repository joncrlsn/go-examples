package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
)

type Article struct {
	Id    int
	Title string
	Body  string
	Date  string
}

func main() {
	fmt.Println("===== sqlite ===== ")
	tempFilename := tempFileName()
	fmt.Printf("temp sqlite db: %s \n", tempFilename)
	db, err := sql.Open("sqlite3", tempFilename)
	check(err, "opening database")

	defer db.Close()
	defer os.Remove(tempFilename)

	_, err = db.Exec("DROP TABLE foo")
	_, err = db.Exec("CREATE TABLE foo (id integer)")
	check(err, "creating table")

	res, err := db.Exec("INSERT INTO foo(id) VALUES(?)", 123)
	check(err, "inserting record")

	affected, _ := res.RowsAffected()
	if affected != 1 {
		log.Fatalf("Expected %d for affected rows, but %d:", 1, affected)
	}

	rows, err := db.Query("SELECT id FROM foo")
	check(err, "selecting records")
	defer rows.Close()

	for rows.Next() {
		var result int
		rows.Scan(&result)
		if result != 123 {
			log.Fatalf("Fetched %q; expected %q", 123, result)
		}
	}
	fmt.Println("sqlite done!")
}

func tempFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), "foo-"+hex.EncodeToString(randBytes)+".db")
}

func check(err error, action string) {
	if err != nil {
		log.Fatalf("Error %s: %v\n", action, err)
	}
}
