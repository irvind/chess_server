package chess

import (
	"fmt"
	"strconv"
	"testing"
)

func tMove(moveStr string) *Move {
	result, _ := MakeMoveFromStr(moveStr)
	return result
}

func tMoveFigure(board *Board, firstPosStr string, secondPosStr string) {
	var firstPosY, secondPosY byte
	fmt.Sscanf(firstPosStr[1:2], "%d", &firstPosY)
	fmt.Sscanf(secondPosStr[1:2], "%d", &secondPosY)

	firstX, firstY, _ := positionToBoardIdx(Position{X: firstPosStr[0], Y: firstPosY})
	secondX, secondY, _ := positionToBoardIdx(Position{X: secondPosStr[0], Y: secondPosY})
	board.Positions[secondY][secondX] = board.Positions[firstY][firstX]
	board.Positions[firstY][firstX] = nil
}

func TestCanMovePawnError(t *testing.T) {
	testCases := []struct {
		Move        *Move
		ExpectedErr error
	}{
		{tMove("e1e4"), ErrInvalidFigure},
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

func TestCanMovePawnResult(t *testing.T) {
	board := NewBoard()
	board2 := NewBoard()
	tMoveFigure(board2, "e2", "e4")
	tMoveFigure(board2, "e7", "e5")

	board3 := NewBoard()
	tMoveFigure(board3, "b1", "c3")
	tMoveFigure(board3, "g8", "f6")

	board4 := NewBoard()

	testCases := []struct {
		Board    *Board
		Move     *Move
		TestName string
		Wanted   bool
	}{
		{board, tMove("e2d3"), "white cannot eat empty position 1", false},
		{board, tMove("e2f3"), "white cannot eat empty position 2", false},
		{board, tMove("e2d2"), "white cannot move horizontaly 1", false},
		{board, tMove("e2f2"), "white cannot move horizontaly 2", false},
		{board, tMove("e2e6"), "white y distance is too big", false},
		{board2, tMove("e4e3"), "white cannot go back 1", false},
		{board2, tMove("e4e2"), "white cannot go back 2", false},
		{board3, tMove("c2c3"), "white path is blocked 1", false},
		{board3, tMove("c2c4"), "white path is blocked 2", false},
		{board, tMove("e7d6"), "black cannot eat empty position 1", false},
		{board, tMove("e7f6"), "black cannot eat empty position 2", false},
		{board, tMove("e7d7"), "black cannot move horizontaly 1", false},
		{board, tMove("e7f7"), "black cannot move horizontaly 2", false},
		{board, tMove("e7e3"), "black y distance is too big", false},
		{board2, tMove("e5e6"), "black cannot go back 1", false},
		{board2, tMove("e5e7"), "black cannot go back 2", false},
		{board3, tMove("f7f6"), "black path is blocked 1", false},
		{board3, tMove("f7f5"), "black path is blocked 2", false},
		{board4, tMove("e2e4"), "white positive 1", true},
		{board4, tMove("e2e3"), "white positive 2", true},
		{board4, tMove("a2a3"), "white positive 3", true},
		{board4, tMove("a2a4"), "white positive 4", true},
		{board4, tMove("e7e5"), "black positive 1", true},
		{board4, tMove("e7e6"), "black positive 2", true},
		{board4, tMove("d7d5"), "black positive 3", true},
		{board4, tMove("d7d6"), "black positive 4", true},
		// TODO: check second double move is not allowed
		// TODO: can take figure
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			result, err := CanMovePawn(tc.Board, *tc.Move)
			if err != nil {
				t.Errorf("Wanted nil err but got '%v' err", err)
			}
			if result != tc.Wanted {
				t.Errorf("Wanted %v result but got %v", tc.Wanted, result)
			}
		})
	}
}

func TestCanMoveKnightResult(t *testing.T) {
	t.Run("Check all knight moves", func(t *testing.T) {
		whiteBoard := NewBoard()
		whiteBoard.Clear()
		whiteBoard.AddFigureStrPos(Figure{Side: WhiteSide, FigureType: KnightFigureType}, "e4")

		blackBoard := NewBoard()
		blackBoard.Clear()
		blackBoard.AddFigureStrPos(Figure{Side: BlackSide, FigureType: KnightFigureType}, "e4")

		validMoves := map[string]bool{
			"c5": true,
			"d6": true,
			"f6": true,
			"g5": true,
			"g3": true,
			"f2": true,
			"d2": true,
			"c3": true,
		}

		for i := 'a'; i <= 'h'; i++ {
			for j := 1; j <= 8; j++ {
				if i == 'e' && j == 4 {
					continue
				}
				secondMove := string(i) + strconv.Itoa(j)
				wantedResult := validMoves[secondMove]

				result, err := CanMoveKnigth(whiteBoard, *tMove("e4" + secondMove))
				if err != nil {
					t.Errorf("Wanted nil err but got '%v' err for move %v white", err, secondMove)
				}
				if result != wantedResult {
					t.Errorf("Wanted %v result but got %v for move %v white", wantedResult, result, secondMove)
				}

				result, err = CanMoveKnigth(blackBoard, *tMove("e4" + secondMove))
				if err != nil {
					t.Errorf("Wanted nil err but got '%v' err for move %v black", err, secondMove)
				}
				if result != wantedResult {
					t.Errorf("Wanted %v result but got %v for move %v black", wantedResult, result, secondMove)
				}
			}
		}
	})

	// board := NewBoard()
	testCases := []struct {
		Board    *Board
		Move     *Move
		TestName string
		Wanted   bool
	}{
		// {board, tMove(""), "", true},
		// TODO: check can take figure
		// TODO: check cannot move if position is not empty
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			result, err := CanMoveKnigth(tc.Board, *tc.Move)
			if err != nil {
				t.Errorf("Wanted nil err but got '%v' err", err)
			}
			if result != tc.Wanted {
				t.Errorf("Wanted %v result but got %v", tc.Wanted, result)
			}
		})
	}
}
