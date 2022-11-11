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
	// TODO: get CreatorWhite field from JSON
	headers, ok = c.Request.Header["Auth-Token"]
	if !ok || len(headers) != 1 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Auth token is required"})
		return
	}

	authToken := headers[0]
	player, err := dao.GetPlayerByAuthSecret(authToken)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	newGameID, err := dao.CreateGame(player.ID)

	c.IndentedJSON(http.StatusOK, gin.H{"id": newGameID})
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