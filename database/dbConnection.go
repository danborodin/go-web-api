package database

import (
	"database/sql"
	"log"
)

//DB -> database variable
var DB *sql.DB

//Connect is the function for connecting to db
func Connect() error {
	var err error

	DB, err = sql.Open("sqlite3", "database/sqlite-database.db")

	if err != nil {
		log.Println(err)
	}

	return err
}

//CreateTaskTable initialize the task table for the first time
func CreateTaskTable(db *sql.DB) {
	db.Query(`CREATE TABLE IF NOT EXISTS tasks (
        "id"	INTEGER NOT NULL UNIQUE,
        "Name"	TEXT NOT NULL,
        "Details"	TEXT,
        "Date"	TEXT,
        "Done"	INTEGER DEFAULT 0,
        PRIMARY KEY("id" AUTOINCREMENT)
)
`)
}
