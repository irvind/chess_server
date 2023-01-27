package chess

import "errors"

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
