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

	if movingFigure.Side != WhiteSide && movingFigure.Side != BlackSide {
		return false, ErrInvalidSide
	}

	isWhiteSide := movingFigure.Side == WhiteSide
	isBlackSide := movingFigure.Side == BlackSide

	if firstPosX == secondPosX {
		validWhiteYPos := firstPosY == (secondPosY+1) || firstPosY == (secondPosY+2)
		validBlackYPos := firstPosY == (secondPosY-1) || firstPosY == (secondPosY-2)

		if isWhiteSide && !validWhiteYPos || isBlackSide && !validBlackYPos {
			return false, nil
		}

		var figure *Figure
		if isWhiteSide && firstPosY == (secondPosY+1) ||
			isBlackSide && firstPosY == (secondPosY-1) {

			figure, err = board.getFigureByPosition(move.Second)
		} else if isWhiteSide && firstPosY == (secondPosY+2) ||
			isBlackSide && firstPosY == (secondPosY-2) {

			checkPos := move.First
			if isWhiteSide {
				checkPos.Y += 1
			} else {
				checkPos.Y -= 1
			}
			figure, err = board.getFigureByPosition(checkPos)
		}
		if err != nil {
			return false, err
		}

		if figure != nil {
			return false, nil
		}
	} else if secondPosX == (firstPosX-1) || secondPosX == (firstPosX+1) {
		validWhiteYPos := secondPosY != firstPosY-1
		validBlackYPos := secondPosY != firstPosY+1

		if isWhiteSide && !validWhiteYPos || isBlackSide && !validBlackYPos ||
			move.FigureTaken == 0 {

			return false, nil
		}

		opponentFigure, err := board.getFigureByPosition(move.Second)
		if err != nil {
			return false, err
		}
		if isWhiteSide && opponentFigure.Side == WhiteSide ||
			isBlackSide && opponentFigure.Side == BlackSide ||
			opponentFigure.FigureType != move.FigureTaken {

			return false, nil
		}
	} else {
		return false, nil
	}

	return true, nil
}
