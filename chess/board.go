package chess

import (
	"errors"
	"fmt"
)

type Board struct {
	Positions [8][8]*Figure

	MoveFigureFx func(Move) error
}

func NewBoard() *Board {
	board := new(Board)

	board.Positions[7][0] = &Figure{WhiteSide, RookFigureType}
	board.Positions[7][1] = &Figure{WhiteSide, KnightFigureType}
	board.Positions[7][2] = &Figure{WhiteSide, BishopFigureType}
	board.Positions[7][3] = &Figure{WhiteSide, QueenFigureType}
	board.Positions[7][4] = &Figure{WhiteSide, KingFigureType}
	board.Positions[7][5] = &Figure{WhiteSide, BishopFigureType}
	board.Positions[7][6] = &Figure{WhiteSide, KnightFigureType}
	board.Positions[7][7] = &Figure{WhiteSide, RookFigureType}

	board.Positions[0][0] = &Figure{BlackSide, RookFigureType}
	board.Positions[0][1] = &Figure{BlackSide, KnightFigureType}
	board.Positions[0][2] = &Figure{BlackSide, BishopFigureType}
	board.Positions[0][3] = &Figure{BlackSide, QueenFigureType}
	board.Positions[0][4] = &Figure{BlackSide, KingFigureType}
	board.Positions[0][5] = &Figure{BlackSide, BishopFigureType}
	board.Positions[0][6] = &Figure{BlackSide, KnightFigureType}
	board.Positions[0][7] = &Figure{BlackSide, RookFigureType}

	for i := 0; i < 8; i++ {
		board.Positions[1][i] = &Figure{BlackSide, PawnFigureType}
		board.Positions[6][i] = &Figure{WhiteSide, PawnFigureType}
	}

	board.Init()

	return board
}

func (board *Board) Init() {
	board.MoveFigureFx = board.moveFigure
}

func (board *Board) ApplyStrMoves(strMoves []string) error {
	for _, strMove := range strMoves {
		move, err := MakeMoveFromStr(strMove)
		if err != nil {
			return err
		}
		board.MoveFigure(*move)
	}

	return nil
}

func (board *Board) MoveFigure(move Move) error {
	return board.MoveFigureFx(move)
}

func (board *Board) Clear() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board.Positions[i][j] = nil
		}
	}
}

func (board *Board) ClearIntPos(x int, y int) {
	// TODO
}

func (board *Board) ClearStrPos(pos string) {
	// TODO
}

func (board *Board) AddFigureIntPos(figure Figure, x int, y int) {
	// TODO
}

func (board *Board) AddFigureStrPos(figure Figure, pos string) {
	// TODO
}

func (board *Board) getFigureByPosition(position Position) (*Figure, error) {
	xCoord, yCoord, err := positionToBoardIdx(position)
	if err != nil {
		return nil, err
	}
	return board.Positions[yCoord][xCoord], nil
}

func (board *Board) moveFigure(move Move) error {
	firstPos1, firstPos2, err := positionToBoardIdx(move.First)
	if err != nil {
		return err
	}
	figure := board.Positions[firstPos1][firstPos2]
	if figure == nil {
		return errors.New("start position is empty")
	}

	switch figure.FigureType {
	case 'p':
		err = board.movePawn(move)
	case 'r':
		// TODO
	case 'k':
		// TODO
	case 'b':
		// TODO
	case 'q':
		// TODO
	case 'g':
		// TODO
	default:
		err = errors.New("unknown figure")
	}

	if err != nil {
		return err
	}

	return nil
}

func (board *Board) movePawn(move Move) error {
	firstPosX, firstPosY, err := positionToBoardIdx(move.First)
	secondPosX, secondPosY, err := positionToBoardIdx(move.Second)
	if err != nil {
		return err
	}

	canMove, err := CanMovePawn(board, move)
	if err != nil {
		return err
	}
	if !canMove {
		return ErrInvalidMove
	}

	board.Positions[secondPosX][secondPosY] = board.Positions[firstPosX][firstPosY]
	board.Positions[firstPosX][firstPosY] = nil

	return nil
}

func (board *Board) PrintPosition(position Position) error {
	pos1, pos2, err := positionToBoardIdx(position)
	if err != nil {
		return err
	}

	figure := board.Positions[pos1][pos2]
	if figure == nil {
		fmt.Printf("%c%d: empty position\n", position.X, position.Y)
	} else {
		fmt.Printf("%c%d: side - %c, figure - %c\n", position.X, position.Y, figure.Side, figure.FigureType)
	}

	return nil
}

// TODO: print board method

func positionToBoardIdx(position Position) (byte, byte, error) {
	if position.X < 'a' || position.X > 'h' || position.Y > 8 {
		return 0, 0, ErrInvalidPosition
	}

	x := position.X - 'a'
	y := 8 - position.Y
	return x, y, nil
}
