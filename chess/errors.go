package chess

import (
	"errors"
)

var (
	ErrInvalidPosition   = errors.New("Invalid position")
	ErrStartAndEndIsSame = errors.New("Start and end position cannot be the same")
	ErrInvalidFigure     = errors.New("Invalid figure")
)
