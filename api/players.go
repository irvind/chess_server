package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func postPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "postPlayers"})
}
