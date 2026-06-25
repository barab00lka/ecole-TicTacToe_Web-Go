package service

import (
	"github.com/google/uuid"
	"tictactoe/internal/domain"
)

type GameService interface {
	ValidateState(uid uuid.UUID, new *domain.Board) error
	MakeAMove(uid uuid.UUID) *domain.GameState
	IsOver(uid uuid.UUID) bool
}
