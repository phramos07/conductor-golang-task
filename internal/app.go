package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/health", healthController).Methods("GET")
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func healthController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
