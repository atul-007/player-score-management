package main

import (
	"log"
	"net/http"

	controller "github.com/atul-007/player-score-management/Controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define the API endpoints using the handlers
	r.HandleFunc("/players", controller.CreatePlayerHandler).Methods("POST")
	r.HandleFunc("/players/{id}", controller.UpdatePlayerHandler).Methods("PUT")
	r.HandleFunc("/players/{id}", controller.DeletePlayerHandler).Methods("DELETE")
	r.HandleFunc("/players", controller.GetAllPlayersHandler).Methods("GET")
	r.HandleFunc("/players/rank/{val}", controller.GetPlayerByRankHandler).Methods("GET")
	r.HandleFunc("/players/random", controller.GetRandomPlayerHandler).Methods("GET")

	// Start the HTTP server on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
