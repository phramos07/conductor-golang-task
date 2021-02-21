// Conductor coding challenge API.
//
// OpenAPI doc for the Conductor coding challenge.
//
// Terms Of Service:
//
//     Schemes: http, https
//	   BasePath: /api
//     Version: 1.0.0
//     Contact: Supun Muthutantri<fakemail@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - APIKey:
//
//     SecurityDefinitions:
//     APIKey:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"conductor/internal/control"
	"conductor/internal/middleware"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	uiPath        = "./static/"
	swaggerPrefix = "/swagger/"
)

// Base router. Will be used for swaggerUI server
var router *mux.Router = mux.NewRouter()

// API Router. Will be used for the API endpoints and all the middlewares
// of logging, in-memory db and auth.
func getAPIRouter() *mux.Router {
	apiRouter := router.PathPrefix("/api").Subrouter()
	return apiRouter
}

// Sets up routes in the API router
func setUpRoutes(apiRouter *mux.Router) {
	control.AddHealthRoutes(apiRouter)
}

// Sets up middlewares in the API router
func setUpMiddlewares(apiRouter *mux.Router) {
	apiRouter.Use(middleware.LoggingMiddleware)
	apiRouter.Use(middleware.AuthorizationMiddleware)
}

// APP's entrypoint
func main() {
	apiRouter := getAPIRouter()

	fs := http.FileServer(http.Dir(uiPath))
	router.PathPrefix(swaggerPrefix).Handler(http.StripPrefix(swaggerPrefix, fs))

	setUpRoutes(apiRouter)
	setUpMiddlewares(apiRouter)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
