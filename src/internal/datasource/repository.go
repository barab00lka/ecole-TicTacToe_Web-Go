package datasource

import (
	"errors"
	"sync"
	"tictactoe/internal/domain"
	"github.com/google/uuid"
)


type GameStore struct {
	m sync.Map // m[UUID]*GameState 
}

func NewRepository() Repository[*domain.GameState] {
	return &GameStore{m: sync.Map{}}
}

// Loads gametate from GameStore Repository. Non-nil error if save does not exist
func (stor *GameStore) Load(id uuid.UUID) *domain.GameState {
    val, ok := stor.m.Load(id)
    if !ok {
        return nil
    }
    game, ok := val.(*domain.GameState)
    if !ok {
        return nil
    }
    return game
}

// this function accepts a valid uuid and pointer to the current gamestate. Returns nil on successful save
func (stor *GameStore) Save(Game *domain.GameState) error {
	if Game == nil {
		return errors.New("no game to save")
	}
	if oldGame := stor.Load(Game.Id); oldGame == nil {
        // New player: store the new *GameState directly
        stor.m.Store(Game.Id, Game)
	} else {
        // Copy the new game state into the existing stored game
        *oldGame = *Game
	} 
    return nil
}
