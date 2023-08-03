package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	helper "github.com/atul-007/player-score-management/Helper"
	models "github.com/atul-007/player-score-management/Models"
	"github.com/gorilla/mux"
)

// Create a new player

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate player attributes

	if player.Name == "" || len(player.Name) > 15 || player.Country == "" || player.Score < 0 {
		http.Error(w, "Invalid player data", http.StatusBadRequest)
		return
	}

	// Save player in the database

	savedPlayer, err := helper.CreatePlayer(&player)
	if err != nil {
		http.Error(w, "Error creating player", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedPlayer)
}

func UpdatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var player models.Player
	err = json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if player.Name == "" || len(player.Name) > 15 || player.Score < 0 {
		http.Error(w, "Invalid player data", http.StatusBadRequest)
		return
	}
	// Update player in the database
	updatedPlayer, err := helper.UpdatePlayer(playerID, &player)
	if err != nil {
		http.Error(w, "Error updating player", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPlayer)
}

// Delete player from the database
func DeletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	err = helper.DeletePlayer(playerID)
	if err != nil {
		http.Error(w, "Error deleting player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllPlayersHandler(w http.ResponseWriter, r *http.Request) {
	// Get all players from the database in descending order
	players, err := helper.GetAllPlayers()
	if err != nil {
		http.Error(w, "Error fetching players", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func GetPlayerByRankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rank, err := strconv.Atoi(vars["val"])
	if err != nil {
		http.Error(w, "Invalid rank value", http.StatusBadRequest)
		return
	}
	// Get player by rank from the database
	player, err := helper.GetPlayerByRank(rank)
	if err != nil {
		http.Error(w, "Error fetching player by rank", http.StatusInternalServerError)
		return
	}

	if player == nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func GetRandomPlayerHandler(w http.ResponseWriter, r *http.Request) {
	//Get players randomly
	player, err := helper.GetRandomPlayer()
	if err != nil {
		http.Error(w, "Error fetching random player", http.StatusInternalServerError)
		return
	}

	if player == nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}
