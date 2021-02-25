package control

import (
	"conductor/internal/facade"
	"conductor/internal/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// An AccountsResponse response model
//
// This is used for returning a response with multiple accounts as body
//
// swagger:response accountsResponse
type AccountsResponse struct {
	// in: body
	Payload []model.Account `json:"accounts"`
}

// The AccountCreatedResponse contains the ID of the created account.
//
// swagger:response accountCreatedResponse
type AccountCreatedResponse struct {
	// in: body
	// required: true
	Account struct {
		// required: true
		ID int64 `json:"ID"`
	} `json:"account"`
}

// PostAccountParams model for adding new Account.
//
// This is used for creating a new Account.
//
// swagger:parameters postAccount
type PostAccountParams struct {
	// in: body
	// required: true
	Account struct {
		// required: true
		Status string `json:"status"`
	} `json:"account"`
}

// GetAccountByIdParams params for adding new Account.
//
// This is used for creating a new Account.
//
// swagger:parameters getAccount deleteAccount
type GetAccountByIdParams struct {
	// in: path
	// required: true
	ID uint64 `json:"id"`
}

// UpdateAccountParams params for updating an Account.
//
// This is used for updating an Account.
//
// swagger:parameters updateAccount
type UpdateAccountParams struct {
	// in: path
	// required: true
	ID uint64 `json:"id"`

	// in: query
	// required: true
	Status string `json:"status"`
}

// swagger:operation GET /accounts accounts getAccounts
// ---
// summary: Retrieves all accounts.
// description: Retrieves all accounts stored in the database.
// responses:
//   '200':
//     "$ref": "#/responses/accountsResponse"
//   '204':
//     description: No accounts found.
//     schema:
//       type: string
func getAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := facade.GetAccounts()
	if len(accounts) == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set(contentTypeHeader, contentTypeJSON)
		err := json.NewEncoder(w).Encode(AccountsResponse{Payload: accounts})
		if err != nil {
			panic(err)
		}
	}
}

// swagger:operation GET /accounts/{id} accounts getAccount
// ---
// summary: Retrieves one account by ID.
// description: Retrieve account from the database that matches given ID.
// responses:
//   '200':
//     description: Account retrieved.
//     schema:
//       "$ref": "#/definitions/account"
//   '204':
//     description: No account found with given ID.
//     schema:
//       type: string
func getAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Panic(model.NewrequestError("Bad ID param type.", model.ErrorBadRequest))
	}

	account := facade.GetAccount(accountID)
	if account != nil {
		w.Header().Set(contentTypeHeader, contentTypeJSON)
		err := json.NewEncoder(w).Encode(account)
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// swagger:operation DELETE /accounts/{id} accounts deleteAccount
// ---
// summary: Deletes one account by ID.
// description: Deletes account from the database that matches given ID.
// responses:
//   '200':
//     description: Account deleted.
//   '204':
//     description: No account found with given ID.
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Panic(model.NewrequestError("Bad ID param type.", model.ErrorBadRequest))
	}

	deleted := facade.DeleteAccount(accountID)
	if !deleted {
		w.WriteHeader(http.StatusNoContent)
	}
}

// swagger:operation PUT /accounts/{id} accounts updateAccount
// ---
// summary: Updates the status of one account by ID.
// description: Updates the status of account from the database that matches given ID.
// responses:
//   '200':
//     description: Account updated.
//   '204':
//     description: No account found with given ID.
func updateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Panic(model.NewrequestError("Bad ID param type.", model.ErrorBadRequest))
	}

	status := r.FormValue("status")
	updated := facade.UpdateAccount(status, accountID)
	if !updated {
		w.WriteHeader(http.StatusNoContent)
	}
}

// swagger:operation POST /accounts accounts postAccount
// ---
// summary: Creates new account.
// description: Creates new account in the database.
// responses:
//   '200':
//     "$ref": "#/responses/accountCreatedResponse"
func postAccount(w http.ResponseWriter, r *http.Request) {
	var account PostAccountParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxJSONSize))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account.Account); err != nil {
		log.Panic(model.NewrequestError("Unprocessable entity", model.ErrorUnprocessableJSON))
	}

	lastID := facade.AddAccount(model.Account{Status: account.Account.Status})
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	err = json.NewEncoder(w).Encode(
		AccountCreatedResponse{
			Account: struct {
				ID int64 "json:\"ID\""
			}{ID: lastID}})
	if err != nil {
		panic(err)
	}
}

// AddAccountsRoutes Adds all routes from account to the router.
func AddAccountsRoutes(r *mux.Router) {
	r.HandleFunc(accountsPath, getAccounts).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(accountsPath, postAccount).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc(accountsByIDPath, getAccount).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc(accountsByIDPath, deleteAccount).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc(accountsByIDPath, updateAccount).Methods(http.MethodPut, http.MethodOptions)
}
