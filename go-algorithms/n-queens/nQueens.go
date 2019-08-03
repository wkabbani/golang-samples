package main

import "fmt"

// Board represents a chess-like board where the queens are to be placed
type Board [][]rune

func makeBoard(n int) Board {
	board := make(Board, n)
	for i := 0; i < n; i++ {
		board[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			board[i][j] = '-'
		}
	}
	return board
}

func printBoard(board Board) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%v ", string(board[i][j]))
		}
		fmt.Printf("\n")
	}
}

func isSafe(board Board, row int, col int) bool {
	// check the column
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// check the diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// check the other diagonal
	for i, j := row, col; i >= 0 && j < len(board[row]); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// otherwise it should be safe
	return true
}

// Prints all the solutions to the nQueen problem
// using backtracking with bounding function
func nQueen(board Board, row int) {
	if row == len(board) {
		fmt.Println("======== Solution ========")
		printBoard(board)
		return
	}

	for j := 0; j < len(board[row]); j++ {
		if isSafe(board, row, j) {
			board[row][j] = 'Q'
			nQueen(board, row+1)
			board[row][j] = '-'
		}
	}
}
