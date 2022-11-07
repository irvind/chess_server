package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/irvind/chess_server/dao"
	"github.com/google/uuid"
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
