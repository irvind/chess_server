package chess

import (
	"testing"
)

func move(moveStr string) *Move {
	result, _ := MakeMoveFromStr(moveStr)
	return result
}

func TestCanMovePawnError(t *testing.T) {
	testCases := []struct {
		Move        *Move
		ExpectedErr error
	}{
		{move("e1e4"), ErrInvalidFigure},
		{&Move{First: Position{X: 0, Y: 9}, Second: Position{X: 'e', Y: 2}}, ErrInvalidPosition},
		{&Move{First: Position{X: 'e', Y: 2}, Second: Position{X: 0, Y: 9}}, ErrInvalidPosition},
		// TODO: opponentFigure is invalid
	}

	board := NewBoard()

	for i, tc := range testCases {
		_, err := CanMovePawn(board, *tc.Move)
		if err != tc.ExpectedErr {
			t.Errorf("Test %d: Wanted '%v' err but got '%v' err", i, tc.ExpectedErr, err)
		}
	}
}

func TestCanMovePawnInvalidMove(t *testing.T) {
	// TODO
}
