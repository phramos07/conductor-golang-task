package facade

import (
	"conductor/internal/dbcontext"
	"conductor/internal/model"
)

// AddAccount adds new account.
func AddAccount(account model.Account) int64 {
	lastID := dbcontext.AddAccount(account)

	return lastID
}

// GetAccounts retrieve all accounts
func GetAccounts() []model.Account {
	accounts := dbcontext.GetAccounts()

	return accounts
}

// GetAccount retrieves account by ID
func GetAccount(accountID int64) *model.Account {
	account := dbcontext.GetAccount(accountID)

	return account
}

// DeleteAccount deletes an account by ID
func DeleteAccount(accountID int64) bool {
	findAccount := dbcontext.GetAccount(accountID)
	if findAccount != nil {
		dbcontext.DeleteAccount(accountID)
		return true
	}

	return false
}

// UpdateAccount deletes an account by ID
func UpdateAccount(status string, accountID int64) bool {
	findAccount := dbcontext.GetAccount(accountID)
	if findAccount != nil {
		if status != "" {
			findAccount.Status = status
		}
		dbcontext.UpdateAccount(*findAccount)
		return true
	}

	return false
}
