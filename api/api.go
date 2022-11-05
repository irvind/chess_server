package api

import "github.com/gin-gonic/gin"

func RunApi() {
	router := gin.Default()

	router.GET("/games", getGames)
	router.POST("/games", postGames)
	router.GET("/games/:id", getGameById)
	router.GET("/games/:id/players", getGamePlayers)
	router.POST("/games/:id/players", postGamePlayers)
	router.GET("/games/:id/moves", getGameMoves)
	router.POST("/games/:id/moves", postGameMoves)

	router.Run("localhost:8080")
}
