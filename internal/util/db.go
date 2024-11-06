package util

import (
	"database/sql"
	"ndy/realworld-gin/internal/config"
	"os"
)

func InitDB() *sql.DB {
	// Capture connection properties
	var db *sql.DB
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", config.MysqlConfig.FormatDSN())
	if err != nil {
		Log.Fatal(err.Error())
		os.Exit(1)
	}
	// Check if the connection is alive
	err = db.Ping()
	if err != nil {
		Log.Fatal(err.Error())
		os.Exit(1)
	}
	// Set the database handle
	return db
}
