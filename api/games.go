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

func postGames(c *gin.Context, context Context) {
	player := context["player"].(*dao.Player)
	newGameID, err := dao.CreateGame(player.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": newGameID})
}

func getGameById(c *gin.Context, context Context) {
	game := context["game"].(*dao.Game)
	c.IndentedJSON(http.StatusOK, game)
}

func getGamePlayers(c *gin.Context, context Context) {
	game := context["game"].(*dao.Game)
	gamePlayers, err := dao.GetPlayersByGameId(int(game.ID))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gamePlayers)
}

func postGamePlayers(c *gin.Context, context Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGamePlayers"})
}

func getGameMoves(c *gin.Context, context Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getGameMoves"})
}

func postGameMoves(c *gin.Context, context Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, gin.H{"message": "postGameMoves"})
}

func requireGameID(c *gin.Context, context Context) bool {
	gameID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id param is invalid"})
		return false
	}

	game, err := dao.GetGame(gameID)
	if game == nil && err == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Game was not found"})
		return false
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return false
	}

	context["game"] = game
	return true
}
