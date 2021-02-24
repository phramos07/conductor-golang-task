package facade

import (
	"conductor/internal/dbcontext"
	"conductor/internal/model"
	"net/http"
)

// AddAccount Adds new account.
func AddAccount(account model.Account) {
	dbcontext.AddAccount(account)
}

// GetAccounts Retrieve all accounts
func GetAccounts() []model.Account {
	accounts := dbcontext.GetAccounts()

	if len(accounts) == 0 {
		panic(model.NewrequestError("No accounts found.", http.StatusNoContent))
	}

	return accounts
}
