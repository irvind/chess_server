package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/irvind/chess_server/dao"
)

func getGames(c *gin.Context) {
	games, err := dao.GetGames()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

func postGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGames"})
}

func getGameById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGameById"})
}

func getGamePlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGamePlayers"})
}

func postGamePlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGamePlayers"})
}

func getGameMoves(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGameMoves"})
}

func postGameMoves(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGameMoves"})
}