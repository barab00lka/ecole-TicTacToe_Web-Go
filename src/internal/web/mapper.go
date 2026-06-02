package web

import (
    "tictactoe/internal/domain"
)

func ToDomainBoard(wBoard GameBoard) domain.Board {
    return domain.Board{Cell: wBoard}
}

func FromDomainBoard(dBoard domain.Board) GameBoard {
    return GameBoard(dBoard.Cell)
}

func ToWebScore(ds domain.Score) Score {
    return Score{
        Pl:   ds.Pl,
        Ai:   ds.Ai,
        Draw: ds.Draw,
    }
}

func FromDomainGameState(d *domain.GameState) GameStateResponse {
	return GameStateResponse{
		Board: d.Board.Cell,
		Over: d.Over,
		Winner: d.Winner,
		Score: ToWebScore(d.Score),
	}
}
