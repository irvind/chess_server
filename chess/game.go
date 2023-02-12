package chess

import (
	"errors"

	"github.com/irvind/chess_server/dao"
)

type Game struct {
	ID           int64
	CreatorID    int64
	OpponentID   *int64
	CreatorWhite *bool
	Finished     bool
	Moves        []Move
	Board        *Board
}

func NewGameFromDaoRow(daoGame *dao.Game) *Game {
	var game Game
	game.ID = daoGame.ID
	game.CreatorID = daoGame.CreatedBy
	if daoGame.Opponent.Valid {
		game.OpponentID = new(int64)
		*game.OpponentID = daoGame.Opponent.Int64
	}
	if daoGame.CreatorWhite.Valid {
		game.CreatorWhite = new(bool)
		*game.CreatorWhite = daoGame.CreatorWhite.Bool
	}
	game.Finished = daoGame.Finished

	return &game
}

func (game *Game) CanAddMove() (bool, error) {
	if game.Finished {
		return false, errors.New("Game is already finished")
	}
	if game.CreatorWhite == nil {
		return false, errors.New("Game initiator color is not set")
	}
	if game.OpponentID == nil {
		return false, errors.New("Opponent is not joined")
	}

	return true, nil
}

func (game *Game) FetchMovesFromDao() error {
	// moves, err
	_, err := dao.GetGameMoves(int(game.ID))
	if err != nil {
		return err
	}

	// TODO: assert move order is correct (check Index)
	// TODO: parse moves to Move structs
	game.Moves = make([]Move, 0)
	return nil
}

func (game *Game) buildBoard() {
	// TODO
}

func (game *Game) ValidateMoveStr(moveStr string, playerID int64) (bool, error) {
	if game.CreatorWhite == nil {
		return false, errors.New("CreatorWhite must be set")
	}
	if game.OpponentID == nil {
		return false, errors.New("OpponentID must be set")
	}

	var validPlayerID int64
	isWhiteMove := len(game.Moves)%2 == 0
	if (*game.CreatorWhite && isWhiteMove) || (!(*game.CreatorWhite) && !isWhiteMove) {
		validPlayerID = game.CreatorID
	} else {
		validPlayerID = *game.OpponentID
	}

	// TODO: parse moveStr, check is board is build, check move

	return validPlayerID == playerID, nil
}

// func NewGameFromDaoMoves(ID int, moves []dao.Move) (*Game, error) {
// 	return nil, nil
// }
