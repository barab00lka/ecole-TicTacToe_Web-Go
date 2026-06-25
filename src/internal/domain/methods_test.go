package domain

import (
	"fmt"
	"testing"
)

var (
	game = GameState{
		Board: Board{
			Cell: [H][W]int8{{'p', 'a', 'p'}, {'a', 'p', 'p'}, {'a', 0, 0}}}, Over: false,
	} // cur gamestate

	game2 = GameState{
		Board: Board{
			Cell: [H][W]int8{{'p', 'a', 'p'}, {'a', 'a', 'p'}, {'a', 0, 0}}}, Over: false,
	}

	empty = GameState{
		Board: Board{
			Cell: [H][W]int8{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}, Over: false,
	}
)

func TestValidate(t *testing.T) {
	new := Board{
		Cell: [H][W]int8{{'a', 'a', 'p'}, {'a', 'p', 'p'}, {'a', 0, 0}},
	} // new. should be invalid

	if err := game.ValidateState(&new); err == nil { // method should throw error
		t.Error("IsValidState didnt notice tampered state")
	} else {
		fmt.Printf("Received error: \"%v\"\n", err)
	}

}

func TestMakeAMoveMaximize(t *testing.T) {
	game2.MakeAMove() // should place 'a' on [2][1] to win
	if game2.Board.Cell[2][1] != Ai {
		t.Error("Ai failed to make the best move")
	} else {
		PrintBoard(game2)
	}

	diag := GameState{
		Board: Board{
			Cell: [H][W]int8{{'a', 'p', 'p'}, {'a', 'a', 'p'}, {'p', 0, 0}}}, Over: false,
	}
	diag.MakeAMove() // should place Ai on [2,2]
	if diag.Board.Cell[2][2] != Ai {
		t.Error("Ai failed to make the best move")
		PrintBoard(diag)
	} else {
		PrintBoard(diag)
	}
}

func TestMakeAMoveMinimize(t *testing.T) {
	game.MakeAMove() // should place 'a' on [2][2] to stop Player from winning
	if game.Board.Cell[2][2] != Ai {
		t.Error("Ai failed to make the best move")
	} else {
		PrintBoard(game)
	}

}

func TestRandomFirstMove(t *testing.T) {
	empty.MakeAMove()
	PrintBoard(empty)
}

func TestGameOver(t *testing.T) {
	board := GameState{
		Board: Board{
			Cell: [H][W]int8{{'p', 'a', 'p'}, {'a', 'p', 'p'}, {'a', 'p', 'a'}}}, Over: false,
	}
	board.Over = board.IsOver() // should be game over because board full and is a draw
	if board.Over != true {
		t.Error("Failed to capture GameOver state")
		PrintBoard(empty)
	}
	board = GameState{
		Board: Board{
			Cell: [H][W]int8{{0, 'a', 'p'}, {'a', 'p', 'p'}, {'a', 'p', 'a'}}}, Over: false,
	}
	board.Over = board.IsOver() // should not be over because cells left
	if board.Over != false {
		t.Error("Failed to capture GameOver state")

	}
}

func TestGameOverWithWin(t *testing.T) {
	board := GameState{
		Board: Board{
			Cell: [H][W]int8{{'a', 'a', 'p'}, {'a', 'p', 'p'}, {'a', 'p', 0}}}, Over: false,
	}
	board.Over = board.IsOver()
	if board.Over != true && board.Winner != Ai {
		t.Error("Failed to capture GameOver state")
		PrintBoard(board)
	} else {
		fmt.Printf("%c won\n", board.Winner)
	}
	board = GameState{
		Board: Board{
			Cell: [H][W]int8{{0, 'a', 'p'}, {'a', 'p', 'p'}, {'p', 'p', 'a'}}}, Over: false,
	}
	board.Over = board.IsOver()
	if board.Over != true && board.Winner != Player {
		t.Error("Failed to capture GameOver state")
		PrintBoard(board)
	} else {
		fmt.Printf("%c won\n", board.Winner)
	}
}
