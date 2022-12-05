package main

import (
	"fmt"

	"github.com/irvind/chess_server/api"
	"github.com/irvind/chess_server/chess"
	"github.com/irvind/chess_server/database"
)

func pointerTest() {
	i, j := 42, 100500

	p := &i
	fmt.Println(*p)
	*p = 100
	fmt.Println(*p)
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(*p)
	fmt.Println(j)
}

func chessBoardTest() {
	position := chess.Position{'a', 1}
	fmt.Printf("%c\n", position.X)

	move := chess.Move{chess.Position{'e', 2}, chess.Position{'e', 4}}
	fmt.Printf("%c%d%c%d\n", move.First.X, move.First.Y, move.Second.X, move.Second.Y)

	var moves []chess.Move
	moves = append(
		moves,
		chess.Move{chess.Position{'e', 2}, chess.Position{'e', 4}},
		chess.Move{chess.Position{'e', 7}, chess.Position{'e', 5}},
		chess.Move{chess.Position{'c', 1}, chess.Position{'d', 3}},
	)
	fmt.Printf("%c%d\n", moves[2].Second.X, moves[2].Second.Y)

	board := chess.NewBoard()

	move2 := chess.Move{chess.Position{'e', 2}, chess.Position{'e', 4}}
	board.PrintPosition(move2.First)
	board.PrintPosition(move2.Second)
	if err := board.MoveFigure(move2); err != nil {
		fmt.Println(err)
	}

	board.PrintPosition(move2.First)
	board.PrintPosition(move2.Second)

	move3, err := chess.NewMove('e', 1, 'd', 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("move3", move3)
}

func main() {
	database.ConnectToDatabase()
	api.RunApi()
}
