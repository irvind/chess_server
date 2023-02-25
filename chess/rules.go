package chess

func CanMovePawn(board *Board, move Move) (bool, error) {
	// TODO: add tests
	firstPosX, firstPosY := positionToBoardIdx(move.First)
	secondPosX, secondPosY := positionToBoardIdx(move.Second)
	figure := board.Positions[firstPosX][firstPosY]

	if figure.Side == WhiteSide {
		if firstPosX == secondPosX {
			if firstPosY != (secondPosY+1) && firstPosY != (secondPosY+2) {
				return false, nil
			}
		} else if secondPosX == (firstPosX-1) || secondPosX == (firstPosX+1) {
			if secondPosY != firstPosY-1 || move.FigureTaken == 0 {
				return false, nil
			}

			opponentFigure, err := board.getFigureByPosition(move.Second)
			if err != nil {
				return false, err
			}
			if opponentFigure.Side == 'w' || opponentFigure.FigureType != move.FigureTaken {
				return false, nil
			}
		} else {
			return false, nil
		}
	} else {
		// TODO: refactor
	}

	return true, nil
}
