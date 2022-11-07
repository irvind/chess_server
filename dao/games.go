package dao

import "github.com/irvind/chess_server/database"

type Game struct {
	ID int64
}

func GetGames() ([]Game, error) {
	var games []Game
	db := database.GetDB()

	rows, err := db.Query("SELECT id FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
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