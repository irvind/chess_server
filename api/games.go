package api

import (
	"net/http"
	"strconv"

	// "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/irvind/chess_server/dao"
)

// type PostGamesParams struct {
// 	CreatorWhite sql.NullBool `json:"creatorIsWhite" binding:"required"`
// }

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
	headers, ok := c.Request.Header["Auth-Token"]
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
	if player == nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Player was not found"})
		return
	}

	newGameID, err := dao.CreateGame(player.ID)

	c.IndentedJSON(http.StatusOK, gin.H{"id": newGameID})
}

func getGameById(c *gin.Context, game *dao.Game) {
	c.IndentedJSON(http.StatusOK, game)
}

func getGamePlayers(c *gin.Context, game *dao.Game) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGamePlayers"})
}

func postGamePlayers(c *gin.Context, game *dao.Game) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGamePlayers"})
}

func getGameMoves(c *gin.Context, game *dao.Game) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGameMoves"})
}

func postGameMoves(c *gin.Context, game *dao.Game) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGameMoves"})
}

func makeGameIdHandler(fn func(*gin.Context, *dao.Game)) gin.HandlerFunc {
	return func(c *gin.Context) {
		gameID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id param is invalid"})
			return
		}
		game, err := dao.GetGame(gameID)

		if game == nil && err == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Game was not found"})
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		fn(c, game)
	}
}
