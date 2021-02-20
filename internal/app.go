package main

import (
	"conductor/internal/control"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router = mux.NewRouter()

func setUpRoutes(apiRouter *mux.Router) {

	control.AddHealthRoutes(apiRouter)
}

func getAPIRouter() *mux.Router {

	apiRouter := router.PathPrefix("/api").Subrouter()
	return apiRouter
}

func main() {

	apiRouter := getAPIRouter()
	//TODO: Add auth middleware to apiRouter

	setUpRoutes(apiRouter)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
