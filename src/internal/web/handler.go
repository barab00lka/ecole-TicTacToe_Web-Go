package web

import (
	"tictactoe/internal/domain"
	"tictactoe/internal/datasource"
	"net/http"
	"errors"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/fatih/color"
)

type Handler struct {
	db datasource.Repository[*domain.GameState]
}

func NewHandler(repo datasource.Repository[*domain.GameState]) *Handler {
	return &Handler{
		db: repo, 
	}
}

func readRequest(r * http.Request) (PlayerReq, error){
	if r.Method != http.MethodPost {
		return PlayerReq{}, errors.New("invalid HTTP method")
	}
    // Get UUID from URL path
    idStr := r.PathValue("current_game_UUID")
    if idStr == "" {
		return PlayerReq{}, errors.New("missing game UUID")
    }
	
	// Parse UUID
    webUUID, err := uuid.Parse(idStr)
    if err != nil {
        return PlayerReq{}, errors.New("invalid UUID format")
    }
	req := PlayerReq{ Id: webUUID }

	// Decode JSON request body 
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return req, errors.New("invalid JSON")
    }

	return req, nil 
}

func (h *Handler) PlayerNewMove(w http.ResponseWriter, r *http.Request) {

	req, err := readRequest(r)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	playerBoard := ToDomainBoard(req.Grid)
	playerUUID := req.Id

	// if save exists
	if game := h.db.Load(playerUUID); game != nil {
		if game.Over {
			// Clear the board
			game.Board = domain.Board{}
			game.Over = false 
			game.Winner = 0 

			c := color.New(color.FgHiYellow)
			c.Printf("[WEB]: uid: %v, msg: Created new Board. Games played: %d\n", playerUUID, game.Score.Ai + game.Score.Pl + game.Score.Draw)
		}
		if err := game.ValidateState(&playerBoard); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

			domain.PrintBoard(*game)

    		return
		} else {
			game.Board = playerBoard // update board with player's move 
			game.MakeAMove() // calculate computer move and update board
			domain.PrintBoard(*game)
			if err := h.db.Save(game); err != nil {
				panic(err) // in case of Repository Save error
			}
			// Craft http json response 
		    resp := FromDomainGameState(game)
		    w.Header().Set("Content-Type", "application/json")
		    w.WriteHeader(http.StatusOK)
		    json.NewEncoder(w).Encode(resp)
		}
	// if no prior save exists
	} else {
		// Init game for player
		newGame := domain.NewGameService(playerUUID)
		c := color.New(color.FgHiYellow)
		c.Printf("[WEB]: uid: %v, msg: Initialized new GameState\n", playerUUID)
		if err := h.db.Save(newGame); err != nil {
			panic(err) // in case of Repository Save error
		}
		// Invalid first move
		if err := newGame.ValidateState(&playerBoard); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    		return
		// Valid first move (player can go first or ai goes first if player sends blank board
		} else {
			newGame.Board = playerBoard // update board with player's move 
			newGame.MakeAMove() // calculate computer move and update board
			domain.PrintBoard(*newGame)
			if err := h.db.Save(newGame); err != nil {
				panic(err) // in case of Repository Save error
			}
			// Craft http json response 
		    resp := FromDomainGameState(newGame)
		    w.Header().Set("Content-Type", "application/json")
		    w.WriteHeader(http.StatusOK)
		    json.NewEncoder(w).Encode(resp)
		}
	}
}
