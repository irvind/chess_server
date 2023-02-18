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

	board.Positions[7][0] = &Figure{'w', 'r'}
	board.Positions[7][1] = &Figure{'w', 'n'}
	board.Positions[7][2] = &Figure{'w', 'b'}
	board.Positions[7][3] = &Figure{'w', 'q'}
	board.Positions[7][4] = &Figure{'w', 'k'}
	board.Positions[7][5] = &Figure{'w', 'b'}
	board.Positions[7][6] = &Figure{'w', 'n'}
	board.Positions[7][7] = &Figure{'w', 'r'}

	board.Positions[0][0] = &Figure{'b', 'r'}
	board.Positions[0][1] = &Figure{'b', 'n'}
	board.Positions[0][2] = &Figure{'b', 'b'}
	board.Positions[0][3] = &Figure{'b', 'q'}
	board.Positions[0][4] = &Figure{'b', 'k'}
	board.Positions[0][5] = &Figure{'b', 'b'}
	board.Positions[0][6] = &Figure{'b', 'n'}
	board.Positions[0][7] = &Figure{'b', 'r'}

	for i := 0; i < 8; i++ {
		board.Positions[1][i] = &Figure{'b', 'p'}
		board.Positions[6][i] = &Figure{'w', 'p'}
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

func (board *Board) getFigureByPosition(position Position) (*Figure, error) {
	// TODO: check for errors
	xCoord, yCoord := positionToBoardIdx(position)
	return board.Positions[xCoord][yCoord], nil
}

func (board *Board) moveFigure(move Move) error {
	firstPos1, firstPos2 := positionToBoardIdx(move.First)
	figure := board.Positions[firstPos1][firstPos2]
	if figure == nil {
		return errors.New("start position is empty")
	}

	var err error
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
	firstPosX, firstPosY := positionToBoardIdx(move.First)
	secondPosX, secondPosY := positionToBoardIdx(move.Second)

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

func (board *Board) PrintPosition(position Position) {
	pos1, pos2 := positionToBoardIdx(position)
	figure := board.Positions[pos1][pos2]
	if figure == nil {
		fmt.Printf("%c%d: empty position\n", position.X, position.Y)
	} else {
		fmt.Printf("%c%d: side - %c, figure - %c\n", position.X, position.Y, figure.Side, figure.FigureType)
	}
}

// TODO: print board method

func positionToBoardIdx(position Position) (byte, byte) {
	// 1 -> 7
	// 8 -> 0
	x := position.X - 'a'
	y := 8 - position.Y
	return y, x
}
