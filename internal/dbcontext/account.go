package dbcontext

import (
	"conductor/internal/model"
	"fmt"
	"log"
)

const (
	createAccountQuery = `CREATE TABLE IF NOT EXISTS Account (
							"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
							"status" VARCHAR,
							"created_at" VARCHAR,
							"updated_at" VARCHAR,
							"deleted_at" VARCHAR		
						);`

	addAccountQuery = `INSERT INTO Account(status, created_at) 
						VALUES (?, datetime('now'));`

	deleteAccountQuery = `UPDATE Account
							SET deleted_at = datetime('now')
							WHERE id = ?;`

	updateAccountQuery = `UPDATE Account
							SET status = ?,
								updated_at = datetime('now')
							WHERE id = ?;`

	getAllAccountsQuery = `SELECT * FROM Account;`

	getAccountQuery = `SELECT * FROM Account 
						WHERE id = ?`
)

// CreateAccountTable Create table Account on DB
func CreateAccountTable() {
	statement, err := GetDbContext().Prepare(createAccountQuery)
	if err != nil {
		log.Fatal("Couldn't create Account table.", err)
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't execute query %s\n", createAccountQuery), err)
	}
}

// AddAccount Add new to the Account table
func AddAccount(account model.Account) int64 {
	statement, err := GetDbContext().Prepare(addAccountQuery)
	if err != nil {
		log.Panic("Couldn't prepare drive for Add account query")
	}
	defer statement.Close()

	result, err := statement.Exec(account.Status)
	if err != nil {
		log.Panic(fmt.Sprintf("Couldn't execute query %s\n", addAccountQuery))
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Panic("Couldn't retrieve last ID, enitity was not added. Rolling back.")
	}

	return lastID
}

// GetAccounts gets all accounts
func GetAccounts() []model.Account {
	row, err := GetDbContext().Query(getAllAccountsQuery)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var accounts []model.Account
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

// DeleteAccount Is called in order to delete an Account entry from the database.
func DeleteAccount(accountID int64) {
	statement, err := GetDbContext().Prepare(deleteAccountQuery)
	if err != nil {
		log.Panic("Couldn't prepare driver for delete account query.", err)
	}
	defer statement.Close()

	_, err = statement.Exec(accountID)
	if err != nil {
		log.Panic(
			fmt.Sprintf(
				"Couldn't execute delete account query %s\n. Rolling back.",
				createAccountQuery,
			), err)
	}
}

// UpdateAccount Is called in order to delete an Account entry from the database.
func UpdateAccount(account model.Account) {
	statement, err := GetDbContext().Prepare(updateAccountQuery)
	if err != nil {
		log.Panic("Couldn't prepare driver for update account query.", err)
	}
	defer statement.Close()

	_, err = statement.Exec(account.Status, account.ID)
	if err != nil {
		log.Panic(
			fmt.Sprintf(
				"Couldn't execute update account query %s\n. Rolling back.",
				createAccountQuery,
			), err)
	}
}

// GetAccount returns an account by ID or nil if there isn't one.
func GetAccount(accountID int64) *model.Account {
	row, err := GetDbContext().Query(getAccountQuery, accountID)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var account model.Account
	if row.Next() {
		err = row.Scan(
			&account.ID,
			&account.Status,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.DeletedAt,
		)
		if err != nil {
			panic(err)
		}

		return &account
	}

	return nil
}
