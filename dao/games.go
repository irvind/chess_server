package dao

import (
	"time"
	"github.com/irvind/chess_server/database"
)

type Game struct {
	ID				int64			`json:"id"`
	CreatedBy		int64			`json:"createdBy"`
	Opponent		JsonNullInt64	`json:"opponent"`
	CreatorWhite	JsonNullBool	`json:"creatorIsWhite"`
	CreatedAt		time.Time		`json:"createdAt"`
}

func GetGames() ([]Game, error) {
	var games []Game
	db := database.GetDB()

	rows, err := db.Query("SELECT id, created_by, opponent, creator_white, created_at FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var game Game
		err = rows.Scan(
			&game.ID,
			&game.CreatedBy,
			&game.Opponent,
			&game.CreatorWhite,
			&game.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func CreateGame() (int64, error) {
	db := database.GetDB()
	row := db.QueryRow("INSERT INTO games (created_at) VALUES (NOW()) RETURNING id")

	var newId int64 = 0
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}
