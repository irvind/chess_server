package chess

import (
	"fmt"
)

type Position struct {
	X byte
	Y byte
}

type Figure struct {
	Side byte
	FigureType byte
}

type Move struct {
	First       Position
	Second      Position
	FigureTaken byte
	Castling    bool
}

func NewMove(startX, startY, endX, endY byte, figureTaken byte, castling bool) (*Move, error) {
	if startX < 'a' || startX > 'h' || startY < 1 || startY > 8 ||
		endX < 'a' || endX > 'h' || endY < 1 || endY > 8 {
		return nil, ErrInvalidPosition
	}

	if startX == endX && startY == endY {
		return nil, ErrStartAndEndIsSame
	}

	switch figureTaken {
	case
		'r',
		'n',
		'b',
		'q',
		'k',
		'p',
		0:
	default:
		return nil, ErrInvalidFigure
	}

	move := Move{
		First:       Position{startX, startY},
		Second:      Position{endX, endY},
		FigureTaken: figureTaken,
		Castling:    castling,
	}
	return &move, nil
}

func MakeMoveFromStr(str string) (*Move, error) {
	if len(str) < 4 {
		return nil, ErrInvalidStrMove
	}

	var startY byte
	_, err := fmt.Sscanf(str[1:2], "%d", &startY)
	if err != nil {
		return nil, ErrInvalidStrMove
	}

	var endY byte
	_, err = fmt.Sscanf(str[3:4], "%d", &endY)
	if err != nil {
		return nil, ErrInvalidStrMove
	}

	var takenFigure byte = 0
	var castling bool = false
	if len(str) == 5 {
		if str[4] == 'c' {
			castling = true
		} else {
			takenFigure = str[4]
		}
	}
	move, err := NewMove(str[0], startY, str[2], endY, takenFigure, castling)
	if err != nil {
		return nil, err
	}

	return move, nil
}
