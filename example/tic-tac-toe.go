package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Position : Position in the board
type Position struct {
	row, col int
}

func main() {
	board := initBoard()

	board = computerMove(board)
	board = computerMove(board)
	board = computerMove(board)
	board = computerMove(board)
	board = computerMove(board)
	board = computerMove(board)
	printBoard(board)

	winner := checkGameOver(board)

	if winner != "NONE" {
		fmt.Printf("GAME OVER, %s WON !!", winner)
	}
}

func printBoard(board [][]string) {
	// Timestmap
	fmt.Printf("\n\n%s\n", time.Now())

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func initBoard() [][]string {
	return [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
}

func playInBoard(pos Position, board [][]string) [][]string {
	updatedBoard := board
	updatedBoard[pos.row][pos.col] = "X"

	return updatedBoard
}

func computerMove(board [][]string) [][]string {
	rand.Seed(time.Now().UnixNano()) // Setting an initial seed for rnd numbers gen

	// Chooses where to play
	row := rand.Intn(3)
	col := rand.Intn(3)

	return playInBoard(Position{row, col}, board)
}

func checkGameOver(board [][]string) string {
	// Check winner, by all the combinations, please refatcor this xD
	switch {
	//// X
	// Horizontal check
	case board[0][0] == "X" && board[0][1] == "X" && board[0][2] == "X":
		return "X"
	case board[0][0] == "X" && board[1][0] == "X" && board[2][0] == "X":
		return "X"
	case board[0][0] == "X" && board[1][0] == "X" && board[2][0] == "X":
		return "X"
	case board[1][0] == "X" && board[1][1] == "X" && board[2][2] == "X":
		return "X"
	case board[2][0] == "X" && board[2][1] == "X" && board[2][2] == "X":
		return "X"
	// Vertical check
	case board[0][0] == "X" && board[1][0] == "X" && board[2][0] == "X":
		return "X"
	case board[0][1] == "X" && board[1][1] == "X" && board[1][2] == "X":
		return "X"
	case board[0][2] == "X" && board[1][2] == "X" && board[2][2] == "X":
		return "X"
	// Oblique check
	case board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X":
		return "X"
	case board[0][2] == "X" && board[1][1] == "X" && board[2][2] == "X":
		return "X"
	case board[2][2] == "X" && board[1][1] == "X" && board[0][0] == "X":
		return "X"

	//// Y
	// Horizontal check
	case board[0][0] == "O" && board[0][1] == "O" && board[0][2] == "O":
		return "Y"
	// Vertical check
	case board[0][0] == "O" && board[1][0] == "O" && board[2][0] == "O":
		return "Y"
	// Oblique check
	case board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O":
		return "Y"
	}

	return "NONE"
}
