package dao

import (
	"time"
	"github.com/irvind/chess_server/database"
)

type Player struct {
	ID			int64		`json:"id"`
	Name		string		`json:"name"`
	AuthSecret	string		`json:"authSecret"`
	CreatedAt	time.Time	`json:"createdAt"`
}

func CreatePlayer(name string, authSecret string) (*Player, error) {
	// TODO: remove code dupliation
	db := database.GetDB()
	row := db.QueryRow(
		"INSERT INTO players (name, auth_secret, created_at) VALUES ($1, $2, NOW()) RETURNING id, created_at",
		name,
		authSecret,
	)

	var newID int64
	var createdAt time.Time
	if err := row.Scan(&newID, &createdAt); err != nil {
		return nil, err
	}

	player := &Player{ID: newID, Name: name, AuthSecret: authSecret, CreatedAt: createdAt}

	return player, nil
}

func GetPlayerByAuthSecret(authSecret string) (*Player, error) {
	// TODO: remove code dupliation
	db := database.GetDB()
	row := db.QueryRow("SELECT id, name, auth_secret, created_at FROM players WHERE auth_secret = $1", authSecret)
	var player Player
	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.AuthSecret,
		&player.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
