package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context map[string]interface{}

type ApiHandlerFunc func(*gin.Context, Context)
type ContextFunc func(*gin.Context, Context) bool

func RunApi() {
	router := gin.Default()

	router.GET("/games", getGames)
	router.POST("/games", makeHandler(postGames, []ContextFunc{requireAuthToken}))
	router.GET("/games/:id", makeHandler(getGameById, []ContextFunc{requireGameID}))
	router.GET("/games/:id/players", makeHandler(getGamePlayers, []ContextFunc{requireGameID}))
	router.POST("/games/:id/players/join", makeHandler(postGamePlayersJoin, []ContextFunc{requireGameID, requireAuthToken}))
	router.GET("/games/:id/moves", makeHandler(getGameMoves, []ContextFunc{requireGameID}))
	router.POST("/games/:id/moves", makeHandler(postGameMoves, []ContextFunc{requireGameID, requireAuthToken}))
	router.POST("/players", postPlayers)

	router.Run("localhost:8080")
}

func makeHandler(fn ApiHandlerFunc, contextFns []ContextFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := make(Context)
		for _, contextFn := range contextFns {
			shouldContinue := contextFn(c, context)
			if !shouldContinue {
				return
			}
		}
		fn(c, context)
	}
}

func JSONIntervalServerError(c *gin.Context) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}

func JSONBadRequestError(c *gin.Context, errorMessage string) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errorMessage})
}

func JSONNotFoundError(c *gin.Context, errorMessage string) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": errorMessage})
}
