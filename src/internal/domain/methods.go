package domain

import (
	"errors"
	"math/rand/v2"
	"fmt" //debug
	"github.com/fatih/color"
)

func NewGameService(id [16]byte) *GameState {
	return &GameState{
		Id: id,
	}
}

/*debug print*/
func PrintBoard(game GameState){
	for i :=0; i < 3; i++ {
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

/* does not mutate GameState, only validates new move against current GameState */
func (cur *GameState) ValidateState(new *Board) error  {
	var NMOVES int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			move := new.Cell[i][j] - cur.Board.Cell[i][j]  
			if move != 0 && move != player  { // if not valid move
				return errors.New("GameState Error: invalid move")
			}
			if move == player {
				NMOVES+=1
			}
		}
	}
	/* If board was empty and new board is also empty : the ai goes first and it's not an error */
	// if NMOVES == 0 {
	// 	return errors.New("GameState Error: make a move!") 
	// } else
	if NMOVES > 1 {
		return errors.New("GameState Error: only 1 move allowed")

	}
	return nil 

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

func (s *Score) Update( winner int8) {
	switch winner {
	case player:
		s.Pl++
	case ai:
		s.Ai++
	case 0:
		s.Draw++
	}
}

// changes GameState: Winner if game is over
func (cur *GameState) IsOver() bool {

	winner := FindWinner(&cur.Board, 3)
	switch winner {
	case player, ai:
		cur.Score.Update(winner)
		return true
	case int8(0):
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

/* returns nil on successful game state change */
func (cur *GameState) MakeAMove() {
	// Check for tie
	if IsBoardFull(&cur.Board) {
		c := color.New(color.FgHiCyan)
		c.Printf("[DOMAIN]: uid: %v, msg: board is full and its a tie!\n", cur.Id)
		cur.Over = cur.IsOver()
		return 
	}

	i, j := GetNextMove(*cur)
	cur.Board.Cell[i][j] = ai
	// Check for wins
	cur.Over = cur.IsOver()
}

func GetNextMove(cur GameState) (uint8, uint8) {
	if cur.Board.IsEmpty() {
		return uint8(rand.UintN(uint(H))), uint8(rand.UintN(uint(W)))
	}

	bestScore := -100000
	bi, bj := -1, -1

	// for each possible game tree branch get best move
	for i:=0; i < H; i++ {
		for j:=0; j < W; j++ {
			if cur.Board.Cell[i][j] == 0 {
				cur.Board.Cell[i][j] = ai

				simulatedScore := Minimax(&cur.Board, 0, player)

				cur.Board.Cell[i][j] = 0 // undo initial move

                /* If the score after this move is 
                 more than the best value, then update 
                 best */ 
                if simulatedScore > bestScore {                
                    bi, bj = i,j
                    bestScore = simulatedScore
				}
			}
		}
	}

	return uint8(bi), uint8(bj)
}

