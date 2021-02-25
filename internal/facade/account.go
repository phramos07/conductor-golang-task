package facade

import (
	"conductor/internal/dbcontext"
	"conductor/internal/model"
)

// AddAccount Adds new account.
func AddAccount(account model.Account) int64 {
	lastID := dbcontext.AddAccount(account)

	return lastID
}

// GetAccounts Retrieve all accounts
func GetAccounts() []model.Account {
	accounts := dbcontext.GetAccounts()

	return accounts
}
