package main

import (
	"golang-backend/configs"
	"golang-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)

	log.Println("Server running on port :8080")
	http.ListenAndServe(":8080", router)
}