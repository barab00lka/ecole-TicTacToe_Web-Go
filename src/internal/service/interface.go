package service

import (
	"tictactoe/internal/domain"
	"github.com/google/uuid"
)

type GameService interface {
	ValidateState(uid uuid.UUID, new *domain.Board) error
	MakeAMove(uid uuid.UUID) *domain.GameState 
	IsOver(uid uuid.UUID) bool
}
