package api

import (
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
	router.POST("/games/:id/players", makeHandler(postGamePlayers, []ContextFunc{requireGameID, requireAuthToken}))
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
