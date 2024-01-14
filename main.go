package main

import (
	"bufio"
	"fmt"
	"os"
)

const boardSize = 8

type ChessBoard [boardSize][boardSize]string

func initializeBoard() ChessBoard {
	board := ChessBoard{}
	// Initialize the board with pieces
	// This is a simplified starting position
	board[0] = [boardSize]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	board[1] = [boardSize]string{"P", "P", "P", "P", "P", "P", "P", "P"}
	for i := 2; i < 6; i++ {
		board[i] = [boardSize]string{" ", " ", " ", " ", " ", " ", " ", " "}
	}
	board[6] = [boardSize]string{"p", "p", "p", "p", "p", "p", "p", "p"}
	board[7] = [boardSize]string{"r", "n", "b", "q", "k", "b", "n", "r"}
	return board
}

func printBoard(board ChessBoard) {
	fmt.Println("  a b c d e f g h")
	for row := 0; row < boardSize; row++ {
		fmt.Printf("%d ", row+1)
		for col := 0; col < boardSize; col++ {
			fmt.Printf("%s ", board[row][col])
		}
		fmt.Println()
	}
}

func main() {
	board := initializeBoard()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printBoard(board)
		fmt.Print("Enter move (e.g., 'e2 to e4'): ")
		scanner.Scan()
		move := scanner.Text()

		if isValidMove(move, board) {
			board = makeMove(move, board)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func isValidMove(move string, board ChessBoard) bool {
	// Implement your move validation logic here
	// Check if the move is valid according to the rules
	return true
}

func makeMove(move string, board ChessBoard) ChessBoard {
	// Implement your move logic here
	// Update the board based on the move
	return board
}
