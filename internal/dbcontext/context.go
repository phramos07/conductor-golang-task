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

// GetDbContext Returns dbcontext object
func GetDbContext() *sql.DB {
	return dbcontext
}

// Create Create DBContext.
func Create() {
	var err error
	if _, err = os.Stat(dbFilepath); os.IsNotExist(err) {
		file, err := os.Create(dbFilepath) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}

	dbcontext, err = sql.Open("sqlite3", dbFilepath) // Open the created SQLite File
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Renew db file. To be called before Create(), if needed.
func Renew() {
	err := os.Remove(dbFilepath) // I delete the file to avoid duplicated records.
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Close Closes DBContext. Should be deferred at App's entrypoint.
func Close() {
	dbcontext.Close()
}

// CreateTables Create all tables.
func CreateTables() {
	CreateAccountTable()
}
