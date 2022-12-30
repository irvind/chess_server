package dao

import (
	"time"

	"github.com/irvind/chess_server/database"
)

type Move struct {
	ID        int64     `json:"id"`
	Index     int       `json:"index"`
	Move      string    `json:"move"`
	CreatedAt time.Time `json:"createdAt"`
}

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
