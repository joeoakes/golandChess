package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	// Split the move into source and destination squares
	parts := strings.Split(move, " to ")
	if len(parts) != 2 {
		return false
	}
	fromSquare, toSquare := parts[0], parts[1]

	// Convert algebraic notation to row and column indices
	fromRow, fromCol := algebraicToIndices(fromSquare)
	toRow, toCol := algebraicToIndices(toSquare)

	// Check if the source and destination squares are valid
	if !isValidSquare(fromRow, fromCol) || !isValidSquare(toRow, toCol) {
		return false
	}

	// Get the piece at the source square
	piece := board[fromRow][fromCol]

	// Implement move validation logic based on the piece type
	switch piece {
	case "P", "p": // Pawn
		return isValidPawnMove(fromRow, fromCol, toRow, toCol, board)
	case "N", "n": // Knight
		// Implement knight move validation logic
		return true
	case "R", "r": // Rook
		// Implement rook move validation logic
		return true
	case "B", "b": // Bishop
		// Implement bishop move validation logic
		return true
	case "Q", "q": // Queen
		// Implement queen move validation logic
		return true
	case "K", "k": // King
		// Implement king move validation logic
		return true
	default:
		return false // Invalid piece
	}
}

func algebraicToIndices(square string) (int, int) {
	if len(square) != 2 {
		return -1, -1
	}
	file := square[0]
	rank := square[1] - '1'
	col := int(file - 'a')
	row := int(boardSize - 1 - rank)
	return row, col
}

func isValidSquare(row, col int) bool {
	return row >= 0 && row < boardSize && col >= 0 && col < boardSize
}

func isValidPawnMove(fromRow, fromCol, toRow, toCol int, board ChessBoard) bool {
	// Implement pawn move validation logic
	// Check for legal pawn moves, captures, and en passant
	return true
}

func makeMove(move string, board ChessBoard) ChessBoard {
	// Split the move into source and destination squares
	parts := strings.Split(move, " to ")
	if len(parts) != 2 {
		return board // Return the unchanged board if the move is invalid
	}
	fromSquare, toSquare := parts[0], parts[1]

	// Convert algebraic notation to row and column indices
	fromRow, fromCol := algebraicToIndices(fromSquare)
	toRow, toCol := algebraicToIndices(toSquare)

	// Check if the source and destination squares are valid
	if !isValidSquare(fromRow, fromCol) || !isValidSquare(toRow, toCol) {
		return board // Return the unchanged board if the move is invalid
	}

	// Copy the current board to avoid modifying the original
	newBoard := board

	// Move the piece from the source square to the destination square
	newBoard[toRow][toCol] = board[fromRow][fromCol]
	newBoard[fromRow][fromCol] = " " // Empty the source square

	return newBoard
}
