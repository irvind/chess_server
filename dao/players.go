package dao

import (
	"database/sql"
	"errors"
	"time"

	"github.com/irvind/chess_server/database"
)

type Player struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	AuthSecret string    `json:"authSecret"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PlayerQueryItem struct {
	Player
	initiator bool
}

type GamePlayers struct {
	Creator  *Player `json:"creator"`
	Opponent *Player `json:"opponent"`
}

func CreatePlayer(name string, authSecret string) (*Player, error) {
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
	db := database.GetDB()
	row := db.QueryRow("SELECT id, name, auth_secret, created_at FROM players WHERE auth_secret = $1", authSecret)
	var player Player
	err := row.Scan(
		&player.ID,
		&player.Name,
		&player.AuthSecret,
		&player.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func GetPlayersByGameId(gameId int) (GamePlayers, error) {
	db := database.GetDB()
	query :=
		"SELECT *, TRUE AS initiator FROM players WHERE id = (SELECT created_by FROM games WHERE id = $1) " +
			"union SELECT *, FALSE AS initiator FROM players WHERE id = (SELECT opponent FROM games WHERE id = $1)"

	rows, err := db.Query(query, gameId)
	if err != nil {
		return GamePlayers{}, err
	}
	defer rows.Close()

	var dbPlayers []PlayerQueryItem
	for rows.Next() {
		var item PlayerQueryItem

		err = rows.Scan(
			&item.Player.ID,
			&item.Player.Name,
			&item.Player.AuthSecret,
			&item.Player.CreatedAt,
			&item.initiator,
		)
		if err != nil {
			return GamePlayers{}, err
		}
		dbPlayers = append(dbPlayers, item)
	}

	if len(dbPlayers) == 0 || len(dbPlayers) > 2 {
		return GamePlayers{nil, nil}, errors.New("Invalid player count")
	} else if len(dbPlayers) == 1 {
		player := dbPlayerToPlayer(&dbPlayers[0])
		return GamePlayers{Creator: player}, nil
	}

	// len(dbPlayers) == 2
	var initiatorIdx, opponentIdx int
	if dbPlayers[0].initiator {
		initiatorIdx, opponentIdx = 0, 1
	} else {
		initiatorIdx, opponentIdx = 1, 0
	}

	return GamePlayers{
		Creator:  dbPlayerToPlayer(&dbPlayers[initiatorIdx]),
		Opponent: dbPlayerToPlayer(&dbPlayers[opponentIdx]),
	}, nil

}

func dbPlayerToPlayer(dbPlayer *PlayerQueryItem) *Player {
	return &Player{
		ID:         dbPlayer.ID,
		Name:       dbPlayer.Name,
		AuthSecret: dbPlayer.AuthSecret,
		CreatedAt:  dbPlayer.CreatedAt,
	}
}
