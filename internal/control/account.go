package control

import (
	"conductor/internal/facade"
	"conductor/internal/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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

// AccountRequest model for adding new Account.
//
// This is used for creating a new Account.
//
// swagger:parameters postAccount
type AccountRequest struct {
	// in: body
	// required: true
	Account struct {
		// required: true
		Status string `json:"status"`
	} `json:"account"`
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

// swagger:operation POST /accounts accounts postAccount
// ---
// summary: Creates new account.
// description: Creates new account in the database.
// responses:
//   '200':
//     "$ref": "#/responses/accountCreatedResponse"
func postAccount(w http.ResponseWriter, r *http.Request) {
	var account AccountRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxJSONSize))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account.Account); err != nil {
		panic(model.NewrequestError("Unprocessable entity", model.ERROR_UNPROCESSABLE_JSON))
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

}
