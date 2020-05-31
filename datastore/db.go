package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./helloworld.db")

	if err != nil {
		return nil, err
	}

	return database, nil;
}

// Creates DB if one does
func SetUpDatabase() {
	database, _ := NewDatabase()

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY , description TEXT)")
	statement.Exec()

	database.Close()
}