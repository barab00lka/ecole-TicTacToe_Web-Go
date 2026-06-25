package service

import (
	"errors"
	"fmt" //debug
	"github.com/fatih/color"
	"github.com/google/uuid"
	"tictactoe/internal/datasource"
	"tictactoe/internal/domain"
)

type TicTacToe struct {
	db datasource.Repository[*domain.GameState]
}

func NewService() *TicTacToe {
	return &TicTacToe{
		db: datasource.NewRepository(),
	}
}

func (s *TicTacToe) IsOver(uid uuid.UUID) bool {
	if cur := s.db.Load(uid); cur != nil {
		return cur.Over
	} else {
		return false
	}
}

func (s *TicTacToe) MakeAMove(uid uuid.UUID) *domain.GameState {
	cur := s.db.Load(uid)
	if cur == nil {
		c := color.New(color.FgHiGreen)
		c.Printf("[DB]: uid: %v, msg: N/A GameState. Send empty board or make first move to start new game!.\n", uid)
		return nil

	}

	// Check for tie
	if domain.IsBoardFull(&cur.Board) {
		c := color.New(color.FgHiCyan)
		c.Printf("[DOMAIN]: uid: %v, msg: board is full and its a tie!\n", cur.Id)
	} else {
		i, j := domain.GetNextMove(*cur)
		cur.Board.Cell[i][j] = domain.Ai
	}

	cur.Over = domain.CheckGameOver(cur)
	s.db.Save(cur)
	return cur
}

func (s *TicTacToe) ValidateState(uid uuid.UUID, new *domain.Board) error {
	game := s.db.Load(uid)
	if game == nil {
		game = domain.NewGame(uid)
	} else if game.Over { // Clear the board
		game.Board = domain.Board{}
		game.Over = false
		game.Winner = 0

		c := color.New(color.FgHiYellow)
		c.Printf("[WEB]: uid: %v, msg: Created new Board. Games played: %d\n", uid, game.Score.Ai+game.Score.Pl+game.Score.Draw)
	}

	if err := s.db.Save(game); err == nil {
		c := color.New(color.FgHiGreen)
		c.Printf("[DB]: uid: %v, msg: Player GameState saved.\n", uid)
	} else {
		panic(err)
	}

	// validation of new game board against existing
	var NMOVES int
	for i := 0; i < domain.H; i++ {
		for j := 0; j < domain.W; j++ {
			move := new.Cell[i][j] - game.Cell[i][j]
			if move != 0 && move != domain.Player { // if not valid move
				return errors.New("GameState Error: invalid move")
			}
			if move == domain.Player {
				NMOVES += 1
			}
		}
	}
	if NMOVES > 1 {
		return errors.New("GameState Error: only 1 move allowed")
	}

	game.Board = *new
	s.db.Save(game)

	return nil
}

/*debug print*/
func PrintBoard(game domain.GameState) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if val := game.Board.Cell[i][j]; val == 0 {
				fmt.Printf("|_")
			} else {
				fmt.Printf("|%c", val)
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Println("+++++++")
}
