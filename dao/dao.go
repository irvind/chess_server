package main

func CreateGame() (int64, error) {
	row := db.QueryRow("INSERT INTO games (created_at) VALUES (NOW()) RETURNING id")

	var newId int64 = 0
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}
