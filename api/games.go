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

type PostGameMovesParams struct {
	Move string `json:"move" binding:"required,max=8"`
}

func getGames(c *gin.Context) {
	games, err := dao.GetGames()
	if err != nil {
		JSONIntervalServerError(c)
		return
	}

	if games == nil {
		c.IndentedJSON(http.StatusOK, []dao.Game{})
		return
	}
	c.IndentedJSON(http.StatusOK, games)
}

func postGames(c *gin.Context, context Context) {
	player := context["player"].(*dao.Player)
	newGameID, err := dao.CreateGame(player.ID)
	if err != nil {
		JSONIntervalServerError(c)
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
		JSONIntervalServerError(c)
		return
	}
	c.IndentedJSON(http.StatusOK, gamePlayers)
}

func postGamePlayersJoin(c *gin.Context, context Context) {
	game := context["game"].(*dao.Game)
	if game.OpponentIsJoined() {
		JSONBadRequestError(c, "Opponent is already joined")
		return
	}

	player := context["player"].(*dao.Player)
	if game.CreatedBy == player.ID {
		JSONBadRequestError(c, "Cannot join the game")
		return
	}

	err := dao.AddPlayerToGame(int(game.ID), int(player.ID))
	if err != nil {
		JSONIntervalServerError(c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

func getGameMoves(c *gin.Context, context Context) {
	game := context["game"].(*dao.Game)
	moves, err := dao.GetGameMoves(int(game.ID))
	if err != nil {
		JSONIntervalServerError(c)
		return
	}

	if moves == nil {
		c.IndentedJSON(http.StatusOK, []dao.Move{})
		return
	}
	c.IndentedJSON(http.StatusOK, moves)
}

func postGameMoves(c *gin.Context, context Context) {
	var params PostGameMovesParams
	if err := c.BindJSON(&params); err != nil {
		// TODO: return error resp
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game := context["game"].(*dao.Game)
	if game.Finished {
		JSONBadRequestError(c, "Game is already finished")
		return
	}
	if !game.CreatorWhite.Valid {
		JSONBadRequestError(c, "Game initiator color is not set")
		return
	}
	if !game.Opponent.Valid {
		JSONBadRequestError(c, "Opponent is not joined")
		return
	}

	moves, err := dao.GetGameMoves(int(game.ID))
	if err != nil {
		JSONIntervalServerError(c)
		return
	}

	var validPlayerID int64
	player := context["player"].(*dao.Player)
	isWhiteMove := len(moves)%2 == 0
	if (game.CreatorWhite.Bool && isWhiteMove) || (!game.CreatorWhite.Bool && !isWhiteMove) {
		validPlayerID = game.CreatedBy
	} else {
		validPlayerID = game.Opponent.Int64
	}
	if validPlayerID != player.ID {
		JSONBadRequestError(c, "It's not your turn")
		return
	}

	// TODO: reconstruct board
	// TODO: check if move is valid

	var moveIndex int
	if moves != nil {
		moveIndex = moves[len(moves)-1].Index + 1
	} else {
		moveIndex = 0
	}
	newMoveID, err := dao.AddMoveToGame(int(game.ID), params.Move, moveIndex)
	if err != nil {
		JSONIntervalServerError(c)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": newMoveID})
}

func requireGameID(c *gin.Context, context Context) bool {
	gameID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		JSONBadRequestError(c, "id param is invalid")
		return false
	}

	game, err := dao.GetGame(gameID)
	if game == nil && err == nil {
		JSONNotFoundError(c, "Game was not found")
		return false
	}
	if err != nil {
		JSONIntervalServerError(c)
		return false
	}

	context["game"] = game
	return true
}
