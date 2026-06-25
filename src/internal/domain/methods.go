package domain

import (
	"math/rand/v2"
)

func NewGame(id [16]byte) *GameState {
	return &GameState{
		Id: id,
	}
}

func IsBoardFull(b *Board) bool {
	for i := range H {
		for j := range W {
			if b.Cell[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (s *Score) Update(winner int8) {
	switch winner {
	case Player:
		s.Pl++
	case Ai:
		s.Ai++
	case 0:
		s.Draw++
	}
}

// changes GameState: Winner if game is over
func CheckGameOver(cur *GameState) bool {
	winner := FindWinner(&cur.Board, 3)
	switch winner {
	case Player, Ai:
		cur.Score.Update(winner)
		return true
	case int8(0): // draw
		if IsBoardFull(&cur.Board) {
			cur.Score.Update(winner) // +1 draw
			return true
		}
	}
	return false
}

func (cur *Board) IsEmpty() bool {
	for i := range H {
		for j := range W {
			if cur.Cell[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func GetNextMove(cur GameState) (uint8, uint8) {
	if cur.Board.IsEmpty() {
		return uint8(rand.UintN(uint(H))), uint8(rand.UintN(uint(W)))
	}

	bestScore := -100000
	bi, bj := -1, -1

	// for each possible game tree branch get best move
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if cur.Board.Cell[i][j] == 0 {
				cur.Board.Cell[i][j] = Ai

				simulatedScore := Minimax(&cur.Board, 0, Player)

				cur.Board.Cell[i][j] = 0 // undo initial move

				/* If the score after this move is
				   more than the best value, then update
				   best */
				if simulatedScore > bestScore {
					bi, bj = i, j
					bestScore = simulatedScore
				}
			}
		}
	}

	return uint8(bi), uint8(bj)
}
