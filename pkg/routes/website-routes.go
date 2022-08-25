package routes

import (
	"github.com/gorilla/mux"
	"github.com/karthik-Prathipati/josh-status-checker/pkg/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/websites", controllers.PublishWebsites).Methods("POST")
	router.HandleFunc("/websites", controllers.GetStatus).Methods("GET")
}
