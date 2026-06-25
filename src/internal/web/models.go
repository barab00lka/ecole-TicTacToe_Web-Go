package web

import "github.com/google/uuid"

type GameBoard [3][3]int8

type PlayerReq struct {
	Grid GameBoard `json:"grid"`
	Id   uuid.UUID
}

type GameStateResponse struct {
	Board  GameBoard `json:"grid"`
	Over   bool      `json:"over"`
	Winner int8      `json:"winner"`
	Score  Score     `json:"score"` // named field (not embedded) for a nested JSON object
}

type Score struct {
	Pl   uint `json:"player_wins"`
	Ai   uint `json:"ai_wins"`
	Draw uint `json:"draws"`
}
