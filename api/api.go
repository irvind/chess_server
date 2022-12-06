package api

import "github.com/gin-gonic/gin"

func RunApi() {
	router := gin.Default()

	router.GET("/games", getGames)
	router.POST("/games", postGames)
	router.GET("/games/:id", makeGameIdHandler(getGameById))
	router.GET("/games/:id/players", makeGameIdHandler(getGamePlayers))
	router.POST("/games/:id/players", makeGameIdHandler(postGamePlayers))
	router.GET("/games/:id/moves", makeGameIdHandler(getGameMoves))
	router.POST("/games/:id/moves", makeGameIdHandler(postGameMoves))
	router.POST("/players", postPlayers)

	router.Run("localhost:8080")
}
