package chess

import (
	"errors"
	"fmt"
)

type Board struct {
	Positions [8][8]*Figure
}

func NewBoard() *Board {
	board := new(Board)

	board.Positions[7][0] = &Figure{'w', 'r'}
	board.Positions[7][1] = &Figure{'w', 'k'}
	board.Positions[7][2] = &Figure{'w', 'b'}
	board.Positions[7][3] = &Figure{'w', 'q'}
	board.Positions[7][4] = &Figure{'w', 'g'}
	board.Positions[7][5] = &Figure{'w', 'b'}
	board.Positions[7][6] = &Figure{'w', 'k'}
	board.Positions[7][7] = &Figure{'w', 'r'}

	board.Positions[0][0] = &Figure{'b', 'r'}
	board.Positions[0][1] = &Figure{'b', 'k'}
	board.Positions[0][2] = &Figure{'b', 'b'}
	board.Positions[0][3] = &Figure{'b', 'q'}
	board.Positions[0][4] = &Figure{'b', 'g'}
	board.Positions[0][5] = &Figure{'b', 'b'}
	board.Positions[0][6] = &Figure{'b', 'k'}
	board.Positions[0][7] = &Figure{'b', 'r'}

	for i := 0; i < 8; i++ {
		board.Positions[1][i] = &Figure{'b', 'p'}
		board.Positions[6][i] = &Figure{'w', 'p'}
	}

	return board
}

func (board *Board) MoveFigure(move Move) error {
	firstPos1, firstPos2 := positionToBoardIdx(move.First)
	if board.Positions[firstPos1][firstPos2] == nil {
		return errors.New("start position is empty")
	}

	// TODO: check by figure

	secondPos1, secondPos2 := positionToBoardIdx(move.Second)
	if board.Positions[secondPos1][secondPos2] != nil {
		return errors.New("end position is not empty")
	}

	board.Positions[secondPos1][secondPos2] = board.Positions[firstPos1][firstPos2]
	board.Positions[firstPos1][firstPos2] = nil

	return nil
}

// TODO: take figure method

func (board *Board) PrintPosition(position Position) {
	pos1, pos2 := positionToBoardIdx(position)
	figure := board.Positions[pos1][pos2]
	if figure == nil {
		fmt.Printf("%c%d: empty position\n", position.X, position.Y)
	} else {
		fmt.Printf("%c%d: side - %c, figure - %c\n", position.X, position.Y, figure.Side, figure.FigureType) 
	}
} 

// TODO: print board

func positionToBoardIdx(position Position) (byte, byte) {
	// 1 -> 7
	// 8 -> 0
	x := position.X - 'a'
	y := 8 - position.Y
	return y, x
}
