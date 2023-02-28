package chess

import (
	"errors"
)

var (
	ErrInvalidStrMove    = errors.New("Invalid move representation")
	ErrInvalidPosition   = errors.New("Invalid position")
	ErrStartAndEndIsSame = errors.New("Start and end position cannot be the same")
	ErrInvalidFigure     = errors.New("Invalid figure")
	ErrInvalidMove       = errors.New("Invalid move")
	ErrInvalidSide       = errors.New("Invalid side")
)
