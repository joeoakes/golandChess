package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const boardSize = 8

type ChessBoard [boardSize][boardSize]string

// Player represents a player in the chess game
type Player struct {
	Name string
}

type ChessGame struct {
	CurrentPlayer *Player
}

// NewPlayer creates a new player with the given name
func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

// SwitchPlayer toggles the current player
func (game *ChessGame) SwitchPlayer() {
	if game.CurrentPlayer.Name == "Player 1" {
		game.CurrentPlayer = NewPlayer("Player 2")
	} else {
		game.CurrentPlayer = NewPlayer("Player 1")
	}
}

// NewChessGame initializes a new chess game
func NewChessGame() *ChessGame {
	return &ChessGame{
		CurrentPlayer: NewPlayer("Player 1"),
	}
}

func initializeBoard() ChessBoard {
	board := ChessBoard{}
	// Initialize the board with pieces
	// This is a simplified starting position
	board[7] = [boardSize]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	board[6] = [boardSize]string{"P", "P", "P", "P", "P", "P", "P", "P"}
	for i := 2; i < 6; i++ {
		board[i] = [boardSize]string{" ", " ", " ", " ", " ", " ", " ", " "}
	}
	board[1] = [boardSize]string{"p", "p", "p", "p", "p", "p", "p", "p"}
	board[0] = [boardSize]string{"r", "n", "b", "q", "k", "b", "n", "r"}
	return board
}

func printBoard(board ChessBoard) {
	fmt.Println("  a b c d e f g h")
	for row := 7; row >= 0; row-- {
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
	chessGame := NewChessGame()

	for {
		fmt.Printf("Current Player: %s\n", chessGame.CurrentPlayer.Name)
		printBoard(board)
		fmt.Print("Enter move (e.g., 'e2 to e4'): ")
		scanner.Scan()
		move := scanner.Text()

		if isValidMove(move, board) {
			board = makeMove(move, board)
			chessGame.SwitchPlayer()
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
		return isValidKnightMove(fromRow, fromCol, toRow, toCol, board)
	case "R", "r": // Rook
		return isValidRookMove(fromRow, fromCol, toRow, toCol, board)
	case "B", "b": // Bishop
		return isValidBishopMove(fromRow, fromCol, toRow, toCol, board)
	case "Q", "q": // Queen
		return isValidQueenMove(fromRow, fromCol, toRow, toCol, board)
	case "K", "k": // King
		//return isValidKingMove(fromRow, fromCol, toRow, toCol, board)
		return true
	default:
		return false // Invalid piece
	}
}

func isValidKnightMove(fromRow, fromCol, toRow, toCol int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= len(board) || toCol < 0 || toCol >= len(board[0]) {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is in an L-shape (2 squares horizontally and 1 square vertically, or vice versa)
	return (rowDiff == 1 && colDiff == 2) || (rowDiff == 2 && colDiff == 1)
}

// isValidRookMove checks if a rook move is valid
func isValidRookMove(fromRow, fromCol, toRow, toCol int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= len(board) || toCol < 0 || toCol >= len(board[0]) {
		return false
	}

	// Check if the move is either vertical or horizontal
	if fromRow == toRow || fromCol == toCol {
		return true
	}

	return false
}

// isValidBishopMove checks if a bishop move is valid
func isValidBishopMove(fromRow, fromCol, toRow, toCol int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= len(board) || toCol < 0 || toCol >= len(board[0]) {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is diagonal (equal absolute differences in rows and columns)
	if rowDiff == colDiff {
		return true
	}

	return false
}

// isValidQueenMove checks if a queen move is valid
func isValidQueenMove(fromRow, fromCol, toRow, toCol int, board ChessBoard) bool {
	// Check if the destination is within the chessboard bounds
	if toRow < 0 || toRow >= len(board) || toCol < 0 || toCol >= len(board[0]) {
		return false
	}

	// Calculate the absolute difference in rows and columns
	rowDiff := abs(fromRow - toRow)
	colDiff := abs(fromCol - toCol)

	// Check if the move is either vertical, horizontal, or diagonal
	return fromRow == toRow || fromCol == toCol || rowDiff == colDiff
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
	// Determine the direction of movement (forward or backward based on player)
	direction := 1 // Assuming white pawns start at the bottom (row 1)
	if board[fromRow][fromCol] == "p" {
		direction = -1 // Black pawns move in the opposite direction
	}

	// Check for a regular pawn move (one square forward)
	if fromCol == toCol && toRow == fromRow+direction && board[toRow][toCol] == " " {
		return true
	}

	// Check for the initial double pawn move (two squares forward)
	if fromCol == toCol && toRow == fromRow+2*direction && fromRow == 1 && board[fromRow+direction][toCol] == " " {
		return true
	}

	// Check for capturing an opponent's piece diagonally
	if abs(toCol-fromCol) == 1 && toRow == fromRow+direction {
		targetPiece := board[toRow][toCol]
		if (direction == 1 && targetPiece >= "A" && targetPiece <= "Z") || (direction == -1 && targetPiece >= "a" && targetPiece <= "z") {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
