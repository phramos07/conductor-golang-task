package facade

import (
	"conductor/internal/dbcontext"
	"conductor/internal/model"
)

// AddAccount Adds new account.
func AddAccount(account model.Account) {
	dbcontext.AddAccount(account)
}

// GetAccounts Retrieve all accounts
func GetAccounts() []model.Account {
	accounts := dbcontext.GetAccounts()

	return accounts
}
