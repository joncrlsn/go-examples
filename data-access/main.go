package main

//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

import (
	"github.com/joncrlsn/go-examples/data-access/svc"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var dbType = "sqlite3"
var dbUrl = "temp.db"
var migrateSqlPath = "./migrations"

func main() {

	guids := svc.NewGuids()

	//
	// Setting up the database
	//
	db, err := setupDb(dbType, dbUrl, migrateSqlPath)
	if err != nil || db == nil {
		log.Println("Error setting up the db", err)
		return ""
	}

	//
	// Examples
	//
	//authSvc := svc.NewAuthSvc(db, guids)

	userSvc := svc.NewUserSvc(db, guids)
	var users []svc.User
	if users, err = userSvc.FindByEmail("jon@example.com"); err != nil {
		log.Println("Error finding user by email", err)
		return
	}

	log.Println(users)
}
