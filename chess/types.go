package chess

const (
	PawnFigureType   = 'p'
	KnightFigureType = 'n'
	BishopFigureType = 'b'
	RookFigureType   = 'r'
	QueenFigureType  = 'q'
	KingFigureType   = 'k'
	WhiteSide        = 'w'
	BlackSide        = 'b'
)

type Position struct {
	X byte
	Y byte
}

type Figure struct {
	Side       byte
	FigureType byte
}
