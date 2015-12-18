//
// Sets up the database connection and ensure the schema is migrated.
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"
	"log"
)

func setupDb(dbType string, dbName string, migrateSqlPath string) (db *sqlx.DB, err error) {
	log.Println("=== Checking schema version ===")
	dbUrl := dbType + "://" + dbName

	// Synchronously migrate the database up if needed.
	allErrors, ok := migrate.ResetSync(dbUrl, migrateSqlPath)
	if !ok {
		log.Println("Error migrating schema", allErrors)
		return nil, nil // Program should stop
	}

	// Get database connection to return
	db, err = sqlx.Connect(dbType, dbName)
	if err != nil {
		log.Println("Error connecting to db", err)
		return nil, err
	}

	// success
	return db, nil
}
