package database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/irvind/chess_server/utils"

)

var db *sql.DB

func ConnectToDatabase() {
	var err error
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			utils.GetenvOrDefault("DBHOST", "127.0.0.1"),
			utils.GetenvOrDefault("DBPORT", "9090"),
			utils.GetenvOrDefault("DBNAME", "chess"),
			utils.GetenvOrDefault("DBUSER", "chess"),
			utils.GetenvOrDefault("DBPASSWORD", "chess"),
		),
	)
		
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
}