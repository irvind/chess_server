package chess

import (
	"fmt"
	"testing"
)

func move(moveStr string) *Move {
	result, _ := MakeMoveFromStr(moveStr)
	return result
}

func TestCanMovePawnError(t *testing.T) {
	testCases := []struct {
		Move        *Move
		ExpectedErr error
	}{
		{move("e1e4"), ErrInvalidFigure},
		{&Move{First: Position{X: 0, Y: 9}, Second: Position{X: 'e', Y: 2}}, ErrInvalidPosition},
		{&Move{First: Position{X: 'e', Y: 2}, Second: Position{X: 0, Y: 9}}, ErrInvalidPosition},
		// TODO: opponentFigure is invalid
	}

	board := NewBoard()

	for i, tc := range testCases {
		_, err := CanMovePawn(board, *tc.Move)
		if err != tc.ExpectedErr {
			t.Errorf("Test %d: Wanted '%v' err but got '%v' err", i, tc.ExpectedErr, err)
		}
	}
}

func TestCanMovePawnInvalidMove(t *testing.T) {
	moveFigure := func(board *Board, firstPosStr string, secondPosStr string) {
		var firstPosY, secondPosY byte
		fmt.Sscanf(firstPosStr[1:2], "%d", &firstPosY)
		fmt.Sscanf(secondPosStr[1:2], "%d", &secondPosY)

		first1, first2, _ := positionToBoardIdx(Position{X: firstPosStr[0], Y: firstPosY})
		second1, second2, _ := positionToBoardIdx(Position{X: secondPosStr[0], Y: secondPosY})
		board.Positions[second1][second2] = board.Positions[first1][first2]
		board.Positions[first1][first2] = nil
	}

	board := NewBoard()
	board2 := NewBoard()
	moveFigure(board2, "e2", "e4")
	moveFigure(board2, "e7", "e5")
	fmt.Println(board2.Positions[4][4])

	board3 := NewBoard()
	moveFigure(board3, "b1", "c3")
	moveFigure(board3, "g8", "f6")

	testCases := []struct {
		Board    *Board
		Move     *Move
		TestName string
	}{
		{board, move("e2d3"), "white cannot eat empty position 1"},
		{board, move("e2d2"), "white cannot move horizontaly 1"},
		{board, move("e2f2"), "white cannot move horizontaly 2"},
		{board, move("e2f3"), "white cannot eat empty position 2"},
		{board, move("e2e6"), "white y distance is too big"},
		{board, move("e7d6"), "black cannot eat empty position 1"},
		{board, move("e7f6"), "black cannot eat empty position 2"},
		{board, move("e7d7"), "black cannot move horizontaly 1"},
		{board, move("e7f7"), "black cannot move horizontaly 2"},
		{board, move("e7e3"), "black y distance is too big"},
		{board2, move("e4e3"), "white cannot go back 1"},
		{board2, move("e4e2"), "white cannot go back 2"},
		{board2, move("e5e6"), "black cannot go back 1"},
		{board2, move("e5e7"), "black cannot go back 2"},
		{board3, move("c2c3"), "white path is blocked 1"},
		{board3, move("c2c4"), "white path is blocked 2"},
		{board3, move("f7f6"), "black path is blocked 1"},
		{board3, move("f7f5"), "black path is blocked 2"},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			result, err := CanMovePawn(tc.Board, *tc.Move)
			if err != nil {
				t.Errorf("Wanted nil err but got '%v' err", err)
			}
			if result != false {
				t.Errorf("Wanted false result but got true")
			}
		})
	}
}
