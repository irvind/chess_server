package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/irvind/chess_server/dao"
)

type PostPlayersParams struct {
	Name string `json:"name" binding:"required,min=2,max=64"`
}

func postPlayers(c *gin.Context) {
	var params PostPlayersParams
	if err := c.BindJSON(&params); err != nil {
		// TODO: return error resp
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	authSecret := uuid.New().String()
	newPlayer, err := dao.CreatePlayer(params.Name, authSecret)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newPlayer)
}

func requireAuthToken(c *gin.Context, context Context) bool {
	headers, ok := c.Request.Header["Auth-Token"]
	if !ok || len(headers) != 1 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Auth token is required"})
		return false
	}

	authToken := headers[0]
	player, err := dao.GetPlayerByAuthSecret(authToken)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return false
	}
	if player == nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Player was not found"})
		return false
	}

	context["player"] = player
	return true
}
