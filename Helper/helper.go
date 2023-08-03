package helper

import (
	"errors"
	"math/rand"
	"sync"

	models "github.com/atul-007/player-score-management/Models"
)

// An in-memory storage for players
var players []models.Player
var playerIDCounter int
var mu sync.RWMutex

// CreatePlayer saves a new player in the database
func CreatePlayer(player *models.Player) (*models.Player, error) {
	mu.Lock()
	defer mu.Unlock()

	playerIDCounter++
	player.ID = playerIDCounter
	players = append(players, *player)
	return player, nil
}

// UpdatePlayer updates the attributes of an existing player in the database
func UpdatePlayer(playerID int, updatedPlayer *models.Player) (*models.Player, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, p := range players {
		if p.ID == playerID {
			// Only update Name and Score
			players[i].Name = updatedPlayer.Name
			players[i].Score = updatedPlayer.Score
			return &players[i], nil
		}
	}

	return nil, errors.New("player not found")
}

// DeletePlayer deletes a player from the database
func DeletePlayer(playerID int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, p := range players {
		if p.ID == playerID {
			// Remove player from slice
			players = append(players[:i], players[i+1:]...)
			return nil
		}
	}

	return errors.New("player not found")
}

// GetAllPlayers fetches all players in descending order based on the score
func GetAllPlayers() ([]models.Player, error) {
	mu.RLock()
	defer mu.RUnlock()

	sortedPlayers := make([]models.Player, len(players))
	copy(sortedPlayers, players)
	// Sort players based on the score in descending order
	sortPlayersByScoreDesc(sortedPlayers)
	return sortedPlayers, nil
}

// GetPlayerByRank fetches a player by their rank
func GetPlayerByRank(rank int) (*models.Player, error) {
	mu.RLock()
	defer mu.RUnlock()

	if rank <= 0 || rank > len(players) {
		return nil, errors.New("invalid rank")
	}

	sortedPlayers := make([]models.Player, len(players))
	copy(sortedPlayers, players)
	// Sort players based on the score in descending order
	sortPlayersByScoreDesc(sortedPlayers)

	return &sortedPlayers[rank-1], nil
}

// GetRandomPlayer fetches a random player from the database
func GetRandomPlayer() (*models.Player, error) {
	mu.RLock()
	defer mu.RUnlock()

	if len(players) == 0 {
		return nil, errors.New("no players found")
	}

	// Pick a random player
	index := rand.Intn(len(players))
	return &players[index], nil
}

// Sort players by score in descending order using bubble sort
func sortPlayersByScoreDesc(players []models.Player) {
	for i := 0; i < len(players)-1; i++ {
		for j := 0; j < len(players)-i-1; j++ {
			if players[j].Score < players[j+1].Score {
				players[j], players[j+1] = players[j+1], players[j]
			}
		}
	}
}
