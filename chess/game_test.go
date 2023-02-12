package chess

import (
	"database/sql"
	"testing"
	"time"

	"github.com/irvind/chess_server/dao"
)

func TestNewGameFromDaoRow(t *testing.T) {
	daoRow := &dao.Game{
		ID:        1,
		CreatedBy: 1,
		Opponent: dao.JsonNullInt64{
			NullInt64: sql.NullInt64{Int64: 0, Valid: false},
		},
		CreatorWhite: dao.JsonNullBool{
			NullBool: sql.NullBool{Bool: true, Valid: true},
		},
		CreatedAt: time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
		Finished:  false,
	}
	game := NewGameFromDaoRow(daoRow)
	if game.ID != 1 ||
		game.CreatorID != 1 ||
		game.OpponentID != nil ||
		*game.CreatorWhite != true ||
		game.Finished != false {

		t.Errorf("game %v is invalid", game)
	}
}
