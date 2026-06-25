package domain

func Minimax(b *Board, depth uint8, turn int8) int {
    /* If Maximizer has won the game return his/her 
     evaluated score */
	winner := FindWinner(b, 3)
    if winner == Ai { 
        return 100 - int(depth)
	}

    /* If Minimizer has won the game return his/her 
     evaluated score */
    if winner == Player {
        return -100 - int(depth)
	}

    /* If there are no more moves and no winner then 
     it is a tie */ 
    if IsBoardFull(b) == true {
        return 0
	}

	c := *b // make a copy of the board 

	if turn == Ai {
		best := -100000
		for i:=0; i < H; i++ {
			for j:=0; j < W; j++ {
				if c.Cell[i][j] == 0 {
					c.Cell[i][j] = Ai

					score := Minimax(&c, depth + 1, Player)
					if best < score {
						best = score
					}

					c.Cell[i][j] = 0
				}
			}
		}
		return best
	} else {
		best := 100000
		for i:=0; i < H; i++ {
			for j:=0; j < W; j++ {
				if c.Cell[i][j] == 0 {
					c.Cell[i][j] = Player

					score := Minimax(&c, depth + 1, Ai)
					if score < best {
						best = score

					}

					c.Cell[i][j] = 0
				}
			}
		}
		return best
	}
}

/* searches for a winning sequence of cells and returns who won, or 0 if no winners were found. (97 == int8('a') for Ai, 112 == int8('p') for Player) */
func FindWinner(b *Board, min int) int8 {
	// Check vertical 
	winner := FindWinnerVert(b, H, W, min)
	if winner == 0 {
		// check diagonal left to right
		winner = FindWinnerDiagMain(b, H, W, min)
	}
	if winner == 0 {
		// Check horizontal 
		winner = FindWinnerHor(b, H, W, min)
	}
	if winner == 0 {
		// check diagonal right to left
		winner = FindWinnerDiagAnti(b, H, W, min)
	}
	return winner
}

func FindWinnerDiagMain(b *Board, N, M, min int) int8 {
    for row := 0; row <= N-min; row++ {
        for col := 0; col <= M-min; col++ {
            startCell := b.Cell[row][col]
            if startCell == 0 {
                continue
            }
            win := true
            for k := 1; k < min; k++ {
                if b.Cell[row+k][col+k] != startCell {
                    win = false
                    break
                }
            }
            if win {
                return startCell
            }
        }
    }
    return 0
}

func FindWinnerDiagAnti(b *Board, N, M, min int) int8 {
    for row := 0; row <= N-min; row++ {
        for col := min - 1; col < M; col++ {
            startCell := b.Cell[row][col]
            if startCell == 0 {
                continue
            }
            win := true
            for k := 1; k < min; k++ {
                if b.Cell[row+k][col-k] != startCell {
                    win = false
                    break
                }
            }
            if win {
                return startCell
            }
        }
    }
    return 0
}
func FindWinnerHor(b *Board, N, M, min int) int8 {
    for row := 0; row < N; row++ {
        for col := 0; col <= M-min; col++ {
            startCell := b.Cell[row][col]
            if startCell == 0 {
                continue
            }
            winSequence := true
            for k := 1; k < min; k++ {
                if b.Cell[row][col+k] != startCell {
                    winSequence = false
                    break
                }
            }
            if winSequence {
                return startCell
            }
        }
    }
    return 0
}

func FindWinnerVert(b *Board, N, M, min int) int8 {
    for col := 0; col < M; col++ {
        for row := 0; row <= N-min; row++ {
            startCell := b.Cell[row][col]
            if startCell == 0 {
                continue
            }
            winSequence := true
            for k := 1; k < min; k++ {
                if b.Cell[row+k][col] != startCell {
                    winSequence = false
                    break
                }
            }
            if winSequence {
                return startCell
            }
        }
    }
    return 0
}

