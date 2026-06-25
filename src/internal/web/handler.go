package web

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"tictactoe/internal/service"
)

type Handler struct {
	s service.GameService
}

func NewHandler() *Handler {
	return &Handler{
		s: service.NewService(),
	}
}

func readRequest(r *http.Request) (PlayerReq, error) {
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
	req := PlayerReq{Id: webUUID}

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
	var response any
	var statusCode int

	if err := h.s.ValidateState(playerUUID, &playerBoard); err != nil {
		statusCode = http.StatusBadRequest
		response = map[string]string{"error": err.Error()}
	} else {
		// calculate computer move and update board
		response = FromDomainGameState(h.s.MakeAMove(playerUUID))
		statusCode = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)

}
