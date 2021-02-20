package control

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	path = "/health"
)

// swagger:operation GET /health getHealth
// ---
// summary: Health check. Return API health status.
// description: If the API is online, an OK will be returned.
// responses:
//   200:
//     description: Health check report. API is (apparently) OK.
//     schema:
//	     type: string
//   500:
//     description: Health check report. Something is wrong.
//     schema:
//       type: string
func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// AddHealthRoutes ...
// Adds routes from path Health to the main API router
func AddHealthRoutes(r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("%s", path), getHealth).Methods("GET")
}
