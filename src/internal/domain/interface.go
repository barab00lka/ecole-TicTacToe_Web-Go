package domain

type GameService interface {
	ValidateState(new *Board) error
	MakeAMove() 
	IsOver() bool
}
