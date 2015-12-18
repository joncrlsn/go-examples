//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"database/sql"
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	dbType = "sqlite3"
	dbName = "temp.db"
)

func main() {

	var err error

	//
	// Ensure we have a database to migrate
	//
	{
		var sqlDb *sql.DB
		sqlDb, err = sql.Open(dbType, dbName)
		if err != nil {
			log.Println("Error opening db", err)
			return
		}
		sqlDb.Close()
	}

	//
	// Ensure database schema is migrated to the latest
	//
	dbUrl := dbType + "://" + dbName
	allErrors, ok := migrate.ResetSync(dbUrl, "./migrations")
	if !ok {
		log.Println("Error migrating schema", allErrors)
		return
	}

}
