package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Board represents the game grid. A 2D array of runes (or integers) will be used.
// 0 represents an empty cell.
type Board [20][10]rune // Adjust dimensions as needed (height x width)

// Piece represents the current falling tetromino.
type Piece struct {
	// Shape definition: 4 rows, 4 columns. A rune value of 0 means empty.
	Shape [4][4]rune
	// Top-left anchor position (row, col) on the board
	Row, Col int
}

// Global tracking for the actively falling piece (simplification for now)
var currentPiece Piece

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("--- Tetris Game Initializing ---")

	// 1. Initialize the game board
	board := Board{}
	printBoard(&board)

	// 2. Spawn the first piece
	spawnPiece(&board)
	printBoard(&board) // Show the board with the spawned piece (requires board update logic, for now just showing spawn)
	// Game Loop placeholder
	fmt.Println("\nSuccessfully spawned the first piece. Next step: Implementing the gravity/drop cycle.")
	/*
	for {
		// ... existing code ...		// 5. Redraw Board and Piece
		//time.Sleep(50 * time.Millisecond)
	}
	*/
}

// printBoard takes a pointer to the board and prints its current state to the console.
func printBoard(b *Board) {
	fmt.Println("\n--- Current Board State ---")	// Print top border
	fmt.Println("+" + string(make([]rune, 10)) + "+")
	for r := 0; r < 20; r++ {
		fmt.Print("|")
		for c := 0; c < 10; c++ {
			// Use a placeholder character if the rune is the zero value (empty)
			if b[r][c] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", b[r][c])
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+" + string(make([]rune, 10)) + "+")
}

// getTetrominoes returns a slice of predefined Piece shapes.
func getTetrominoes() []Piece {
	// Define the I piece (Cyan) - Line shape
	i := Piece{
		Shape: [4][4]rune{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Row: 0, Col: 0,
	}

	// Define the O piece (Yellow) - Square shape
	o := Piece{
		Shape: [4][4]rune{
			{1, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Row: 0, Col: 0,
	}

	// Define the T piece (Purple)
	t := Piece{
		Shape: [4][4]rune{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Row: 0, Col: 0,
	}

	// Define the L piece (Orange)
	l := Piece{
		Shape: [4][4]rune{
			{0, 0, 0, 1},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Row: 0, Col: 0,
	}

	return []Piece{i, o, t, l}
}

// spawnPiece selects a random piece type and initializes the currentPiece global variable.
func spawnPiece(board *Board) {
	availablePieces := getTetrominoes()

	// Select a random index
	randIndex := rand.Intn(len(availablePieces))
	pieceToSpawn := availablePieces[randIndex]

	// Reset the global piece tracker
	currentPiece = Piece{
		Shape: pieceToSpawn.Shape,
		Row: 0,
		Col: 0,
	}
	fmt.Printf("\n[INFO] Piece spawned successfully! (Type index: %d)\n", randIndex)
}

