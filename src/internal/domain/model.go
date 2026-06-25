package domain

import "github.com/google/uuid"

const (
	H = 3
	W = 3
)

const (
	Player = int8('p')
	Ai = int8('a')
)

type Board struct {
	Cell [H][W]int8 // 97 == int8('a') for ai, 112 == int8('p') for player, 0 for empty cell
}

type Score struct {
	Pl uint
	Ai uint
	Draw uint
}

type GameState struct {
	Board
	Over bool
	Winner int8
	Score // { player wins, ai wins, draws }
	Id uuid.UUID 
}
