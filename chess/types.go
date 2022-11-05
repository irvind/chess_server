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
	First Position
	Second Position
}

func NewMove(startX, startY, endX, endY byte) (*Move, error) {
	if startX < 'a' || startX > 'h' || startY < 1 || startY > 8 ||
			endX < 'a' || endX > 'h' || endY < 1 || endY > 8 {
		return nil, errors.New("Invalid coords")
	}

	if startX == endX && startY == endY {
		return nil, errors.New("Start and end position cannot be the same")
	}

	move := Move{Position{startX, startY}, Position{endX, endY}}
	return &move, nil
}
