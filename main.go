package main

import (
	"github.com/irvind/chess_server/api"
	"github.com/irvind/chess_server/database"
)

func main() {
	database.ConnectToDatabase()
	api.RunApi()
}
