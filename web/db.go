package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
	"log"
)

func setupDb(dbType string, dbName string, migrateSqlPath string) (db *sqlx.DB, err error) {
	log.Println("=== Checking database version ===")
	dbUrl := dbType + "://" + dbName

	// Synchronously migrate the database up if needed.
	allErrors, ok := migrate.ResetSync(dbUrl, migrateSqlPath)
	if !ok {
		log.Println("Error migrating database", allErrors)
		return // Program should stop
	}

	// Get database connection to return
	db, err = sqlx.Connect(dbType, dbName)
	if err != nil {
		log.Println("Error connecting to db", err)
		return
	}

	// success
	return
}
