package chess

import (
	"testing"
)

func TestMakeMoveFromStr(t *testing.T) {
	invalidPosTestCases := []struct {
		StartX, StartY, EndX, EndY byte
		FigureTaken                byte
		Castling                   bool
		ExpectedErr                error
	}{
		{'Z', 1, 'b', 2, 0, false, ErrInvalidPosition},
		{'a', 1, 'B', 2, 0, false, ErrInvalidPosition},
		{'e', 0, 'e', 4, 0, false, ErrInvalidPosition},
		{'e', 0, 'e', 4, 0, false, ErrInvalidPosition},
		{'e', 1, 'e', 10, 0, false, ErrInvalidPosition},
		{'e', 2, 'e', 2, 0, false, ErrStartAndEndIsSame},
		{'d', 3, 'd', 3, 0, false, ErrStartAndEndIsSame},
		{'e', 2, 'e', 4, 1, false, ErrInvalidFigure},
		{'e', 2, 'e', 4, 'Z', false, ErrInvalidFigure},
	}

	for i, tc := range invalidPosTestCases {
		_, err := NewMove(tc.StartX, tc.StartY, tc.EndX, tc.EndY, tc.FigureTaken, tc.Castling)
		if err != tc.ExpectedErr {
			t.Errorf("Test %d: Wanted '%v' err but got '%v'", i, tc.ExpectedErr, err)
		}
	}

	move, err := NewMove('e', 2, 'e', 4, 0, false)
	if err != nil {
		t.Errorf("Expected nil err")
	}
	expected := Move{
		First:       Position{X: 'e', Y: 2},
		Second:      Position{X: 'e', Y: 4},
		FigureTaken: 0,
		Castling:    false,
	}

	if *move != expected {
		t.Errorf("Wanted %v but got %v", *move, expected)
	}
}
