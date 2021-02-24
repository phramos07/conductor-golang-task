package dbcontext

import (
	"database/sql"
	"log"
	"os"
)

const (
	dbFilepath = "./internal/model/database.db"
)

var dbcontext *sql.DB

// Create ...
func Create() {
	if _, err := os.Stat(dbFilepath); os.IsNotExist(err) {
		file, err := os.Create(dbFilepath) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}

	dbcontext, _ = sql.Open("sqlite3", dbFilepath) // Open the created SQLite File
}

// Renew db file. To be called before Create(), if needed.
func Renew() {
	err := os.Remove(dbFilepath) // I delete the file to avoid duplicated records.
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Close ...
func Close() {
	dbcontext.Close()
}

// CreateTables ...
func CreateTables() {
	// Empty
	// TODO: Call every entity createtable method available in the model package
}
