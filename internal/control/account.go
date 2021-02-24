package control

import (
	"conductor/internal/facade"
	"conductor/internal/model"
	"encoding/json"
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

// swagger:route GET /accounts accounts getAccounts
// ---
// summary: Retrieves all accounts.
// description: Retrieves all accounts stored in the database.
// responses:
//   200: accountsResponse
func getAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := facade.GetAccounts()
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	err := json.NewEncoder(w).Encode(AccountsResponse{Payload: accounts})
	if err != nil {
		panic(err)
	}
}

// AddAccountsRoutes Adds all routes from account to the router.
func AddAccountsRoutes(r *mux.Router) {
	r.HandleFunc(accountsPath, getAccounts).Methods(http.MethodGet, http.MethodOptions)
}
