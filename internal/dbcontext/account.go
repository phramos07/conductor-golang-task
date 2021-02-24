package dbcontext

import (
	"conductor/internal/model"
	"fmt"
	"log"
)

const (
	createQuery = `CREATE TABLE IF NOT EXISTS Account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"status" VARCHAR,
		"created_at" VARCHAR,
		"updated_at" VARCHAR,
		"deleted_at" VARCHAR		
	  );`

	addQuery = `INSERT INTO Account(status, created_at) VALUES (?, datetime('now'))`

	getAllQuery = `SELECT * FROM Account`
)

// CreateAccountTable Create table Account on DB
func CreateAccountTable() {
	statement, err := GetDbContext().Prepare(createQuery)
	if err != nil {
		log.Fatal("Couldn't create Account table.", err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't execute query %s\n", createQuery), err)
	}
}

// AddAccount Add new to the Account table
func AddAccount(account model.Account) {
	statement, err := GetDbContext().Prepare(addQuery)
	if err != nil {
		panic("Couldn't execute Add account query")
	}

	_, err = statement.Exec(account.Status)
	if err != nil {
		panic(fmt.Sprintf("Couldn't execute query %s\n", addQuery))
	}
}

// GetAccounts gets all accounts
func GetAccounts() []model.Account {
	var accounts []model.Account
	row, err := GetDbContext().Query(getAllQuery)
	if err != nil {
		panic(err)
	}

	defer row.Close()
	for row.Next() {
		var tmpAccount model.Account // Iterate and fetch the records from result cursor
		err = row.Scan(
			&tmpAccount.ID,
			&tmpAccount.Status,
			&tmpAccount.CreatedAt,
			&tmpAccount.UpdatedAt,
			&tmpAccount.DeletedAt,
		)
		if err != nil {
			panic(err)
		}

		accounts = append(accounts, tmpAccount)
	}

	return accounts
}
