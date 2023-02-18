package chess

import (
	"testing"
)

type moveStore struct {
	SavedMoves []Move
}

func (store *moveStore) MoveFigure(move Move) error {
	store.SavedMoves = append(store.SavedMoves, move)
	return nil
}

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board.Positions[0][3].FigureType != 'q' {
		t.Errorf("Wanted 'q' but got '%v'", board.Positions[0][3].FigureType)
	}
	if board.Positions[0][3].Side != 'b' {
		t.Errorf("Wanted 'b' but got '%v'", board.Positions[0][3].Side)
	}
}

func TestApplyStrMoves(t *testing.T) {
	board := NewBoard()

	boardMoveStore := &moveStore{}
	board.MoveFigureFx = boardMoveStore.MoveFigure
	err := board.ApplyStrMoves([]string{"e2e4", "a2a4", "d2d4"})
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if len(boardMoveStore.SavedMoves) != 3 {
		t.Errorf("len(SavedMoves) wanted 3 got %v", len(boardMoveStore.SavedMoves))
	}
}
