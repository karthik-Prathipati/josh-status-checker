package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthik-Prathipati/josh-status-checker/pkg/routes"
)

func main() {
	mux := mux.NewRouter()
	routes.RegisterRoutes(mux)
	http.Handle("/", mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
