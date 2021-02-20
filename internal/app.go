package main

import (
	"conductor/internal/control"
	"conductor/internal/middleware"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router = mux.NewRouter()

func getAPIRouter() *mux.Router {
	apiRouter := router.PathPrefix("/api").Subrouter()
	return apiRouter
}

func setUpRoutes(apiRouter *mux.Router) {
	control.AddHealthRoutes(apiRouter)
}

func setUpMiddlewares(apiRouter *mux.Router) {
	apiRouter.Use(middleware.LoggingMiddleware)
}

func main() {
	apiRouter := getAPIRouter()
	//TODO: Add auth middleware to apiRouter

	setUpRoutes(apiRouter)
	setUpMiddlewares(apiRouter)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
