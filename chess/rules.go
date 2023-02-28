package chess

func CanMovePawn(board *Board, move Move) (bool, error) {
	movingFigure, err := board.getFigureByPosition(move.First)
	if err != nil {
		return false, err
	}
	if movingFigure.FigureType != PawnFigureType {
		return false, ErrInvalidFigure
	}

	firstPosX, firstPosY, err := positionToBoardIdx(move.First)
	if err != nil {
		return false, err
	}
	secondPosX, secondPosY, err := positionToBoardIdx(move.Second)
	if err != nil {
		return false, err
	}

	if movingFigure.Side == WhiteSide {
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
