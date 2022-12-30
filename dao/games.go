package dao

import (
	"database/sql"
	"time"

	"github.com/irvind/chess_server/database"
)

type Game struct {
	ID           int64         `json:"id"`
	CreatedBy    int64         `json:"createdBy"`
	Opponent     JsonNullInt64 `json:"opponent"`
	CreatorWhite JsonNullBool  `json:"creatorIsWhite"`
	CreatedAt    time.Time     `json:"createdAt"`
}

func (game *Game) OpponentIsJoined() bool {
	return game.Opponent.Valid
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

func GetGame(id int) (*Game, error) {
	var game Game
	db := database.GetDB()

	row := db.QueryRow(
		"SELECT id, created_by, opponent, creator_white, created_at "+
			"FROM games "+
			"WHERE id = $1",
		id,
	)

	err := row.Scan(
		&game.ID,
		&game.CreatedBy,
		&game.Opponent,
		&game.CreatorWhite,
		&game.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func CreateGame(createdBy int64) (int64, error) {
	db := database.GetDB()
	row := db.QueryRow(
		"INSERT INTO games (created_by, created_at) VALUES ($1, NOW()) RETURNING id",
		createdBy,
	)

	var newId int64 = 0
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}

func AddPlayerToGame(gameID int, playerID int) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE games SET opponent = $1 WHERE id = $2", playerID, gameID)
	if err != nil {
		return err
	}

	return nil
}
