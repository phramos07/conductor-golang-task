// Conductor coding challenge API.
//
// OpenAPI doc for the Conductor coding challenge.
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
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
	apiRouter.Use(middleware.AuthorizationMiddleware)
}

func main() {
	apiRouter := getAPIRouter()
	//TODO: Add auth middleware to apiRouter

	setUpRoutes(apiRouter)
	setUpMiddlewares(apiRouter)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
