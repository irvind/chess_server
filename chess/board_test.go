package chess

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board.Positions[0][3].FigureType != 'q' {
		t.Errorf("Wanted 'q' but got '%v'", board.Positions[0][3].FigureType)
	}
	if board.Positions[0][3].Side != 'b' {
		t.Errorf("Wanted 'b' but got '%v'", board.Positions[0][3].Side)
	}
}
