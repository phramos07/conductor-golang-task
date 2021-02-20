package control

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	path          = "health"
	getHealthPath = ""
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// AddHealthRoutes ...
// Adds routes from path Health to the main API router
func AddHealthRoutes(r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/%s", path, getHealthPath), getHealth).Methods("GET")
}
