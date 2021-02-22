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
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	uiPath          = "./static/"
	swaggerPrefix   = "/swagger/"
	portEnvVariable = "PORT"
	standardPort    = "8080"
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
	apiRouter.Use(mux.CORSMethodMiddleware(apiRouter),
		middleware.CorsMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthorizationMiddleware)

}

// Loads env variables for local development
func loadEnv() {
	log.Println("Loading environment variables for local development")
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}
}

// APP's entrypoint
func main() {
	loadEnv()

	apiRouter := getAPIRouter()

	fs := http.FileServer(http.Dir(uiPath))
	router.PathPrefix(swaggerPrefix).Handler(http.StripPrefix(swaggerPrefix, fs))

	setUpRoutes(apiRouter)
	setUpMiddlewares(apiRouter)

	// $PORT is defined in the server
	var port string
	port, found := os.LookupEnv(portEnvVariable)

	if !found || port == "" {
		port = standardPort
	}

	log.Printf("Listening on localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
