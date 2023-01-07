package dao

import (
	"time"

	"github.com/irvind/chess_server/database"
)

// Move holds move table row from the database
type Move struct {
	ID        int64     `json:"id"`
	Index     int       `json:"index"`
	Move      string    `json:"move"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetGameMoves retrieves all game moves from the database
func GetGameMoves(gameID int) ([]Move, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT id, idx, description, created_at FROM moves WHERE game_id = $1 ORDER BY idx", gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []Move
	for rows.Next() {
		var move Move
		err = rows.Scan(
			&move.ID,
			&move.Index,
			&move.Move,
			&move.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}

// AddMoveToGame adds new move to existing game by game id
func AddMoveToGame(gameID int, move string, index int) (int64, error) {
	db := database.GetDB()

	row := db.QueryRow(
		"INSERT INTO moves(game_id, idx, description, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id",
		gameID,
		index,
		move,
	)

	var newId int64 = 0
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}
