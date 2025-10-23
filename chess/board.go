package chess

import (
	"fmt"
	"math/bits"
)

type Bitboard uint64

const (
	WhitePawn   int = iota // 0
	WhiteKnight            // 1
	WhiteBishop            // 2
	WhiteRook              // 3
	WhiteQueen             // 4
	WhiteKing              // 5
	BlackPawn              // 6
	BlackKnight            // 7
	BlackBishop            // 8
	BlackRook              // 9
	BlackQueen             // 10
	BlackKing              // 11
	NoPiece                // 12 (Useful sentinel value)
)

// Square constants (0 = A1, …, 63 = H8)
const (
	A1 uint = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NoSquare
)

// Board holds the complete state of game using bitboards
type Board struct {
	// Piece Bitboards: One for each piece type and color.
	// index is WhitePawn, WhiteKnight, ..., BlackKing
	PieceBB [12]uint64

	// Occupancy Bitboards: Combined boards for faster lookups.
	// index 0: White pieces, 1: Black pieces, 2: All pieces
	OccupancyBB [3]uint64

	// Game State Variables (add these later)
	SideToMove   int     // e.g., White or Black constants
	Castling     uint8   // Bitflags for castling rights (WK, WQ, BK, BQ)
	EnPassantSq  uint    // The square index (0-63) or NoSquare
	HalfMoveClock int     // For the 50-move rule
	FullMoveNum  int     // Game move counter
}

func NewBoard()*Board {
	b := &Board{}
	b.SideToMove = 0 // 0 for white, 1 for black
	b.Castling = 0b1111 // Example: WK | WQ | BK | BQ
	b.EnPassantSq = NoSquare
	b.HalfMoveClock = 0
	b.FullMoveNum = 1

	// Setup white pieces
	b.PieceBB[WhitePawn] = 0b0000000000000000000000000000000000000000000000001111111100000000
	b.PieceBB[WhiteKnight] = 0b0000000000000000000000000000000000000000000000000000000001000010
	b.PieceBB[WhiteBishop] = 0b0000000000000000000000000000000000000000000000000000000000100100
	b.PieceBB[WhiteRook] = 0b0000000000000000000000000000000000000000000000000000000010000001
	b.PieceBB[WhiteQueen] = 0b0000000000000000000000000000000000000000000000000000000000001000
	b.PieceBB[WhiteKing] = 0b0000000000000000000000000000000000000000000000000000000000010000

	// Setup black pieces
	b.PieceBB[BlackPawn] = 0b0000000011111111000000000000000000000000000000000000000000000000
	b.PieceBB[BlackKnight] = 0b0100001000000000000000000000000000000000000000000000000000000000
	b.PieceBB[BlackBishop] = 0b0010010000000000000000000000000000000000000000000000000000000000
	b.PieceBB[BlackRook] = 0b1000000100000000000000000000000000000000000000000000000000000000
	b.PieceBB[BlackQueen] = 0b0000100000000000000000000000000000000000000000000000000000000000
	b.PieceBB[BlackKing] = 0b0001000000000000000000000000000000000000000000000000000000000000

		// Calculate initial occupancy
    for p := WhitePawn; p <= WhiteKing; p++ {
        b.OccupancyBB[0] |= b.PieceBB[p] // White occupancy
    }
    for p := BlackPawn; p <= BlackKing; p++ {
        b.OccupancyBB[1] |= b.PieceBB[p] // Black occupancy
    }
    b.OccupancyBB[2] = b.OccupancyBB[0] | b.OccupancyBB[1] // All pieces occupancy

	return b
}

// setBit turns the bit 'on' at the given square index (0-63).
func setBit(bb uint64, square uint) uint64 {
	return bb | (1 << square)
}

// clearBit turns the bit 'off' at the given square index.
func clearBit(bb uint64, square uint) uint64 {
	return bb &^ (1 << square) 
}

// getBit checks if the bit is 'on' at the given square index.
func getBit(bb uint64, square uint) bool {
	return (bb & (1 << square)) != 0
}

// popBit finds the index of the least significant bit, clears it from the bitboard pointer,
// and returns the index. Returns NoSquare (64) if bitboard is empty.
func popBit(bb *uint64) uint {
	if *bb == 0 {
		return NoSquare
	}
	index := uint(bits.TrailingZeros64(*bb))
	*bb &^= (1 << index) // Clear the LSB using AND NOT
	return index
}

// Helper methods on your Board struct to manage pieces:
func (b *Board) AddPiece(piece int, square uint) {
	b.PieceBB[piece] = setBit(b.PieceBB[piece], square)
	color := piece / 6 // 0 for white, 1 for black
	b.OccupancyBB[color] = setBit(b.OccupancyBB[color], square)
	b.OccupancyBB[2] = setBit(b.OccupancyBB[2], square)
}

func (b *Board) RemovePiece(piece int, square uint) {
	b.PieceBB[piece] = clearBit(b.PieceBB[piece], square)
	color := piece / 6 // 0 for white, 1 for black
	b.OccupancyBB[color] = clearBit(b.OccupancyBB[color], square)
	b.OccupancyBB[2] = clearBit(b.OccupancyBB[2], square)
}

func (b *Board) GetPieceOnSquare(square uint) int {
	for p := WhitePawn; p <= BlackKing; p++ {
		if getBit(b.PieceBB[p], square) {
			return p
		}
	}
	return NoPiece
}

// PrintBoard displays the board state to the console.
func (b *Board) PrintBoard() {
	fmt.Println("\n  +-----------------+")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d | ", rank+1)
		for file := 0; file < 8; file++ {
			square := uint(rank*8 + file)
			piece := b.GetPieceOnSquare(square)
			
			pieceChar := '.' // Default empty square
			switch piece {
			case WhitePawn:
				pieceChar = 'P'
			case WhiteKnight:
				pieceChar = 'N'
			case WhiteBishop:
				pieceChar = 'B'
			case WhiteRook:
				pieceChar = 'R'
			case WhiteQueen:
				pieceChar = 'Q'
			case WhiteKing:
				pieceChar = 'K'
			case BlackPawn:
				pieceChar = 'p'
			case BlackKnight:
				pieceChar = 'n'
			case BlackBishop:
				pieceChar = 'b'
			case BlackRook:
				pieceChar = 'r'
			case BlackQueen:
				pieceChar = 'q'
			case BlackKing:
				pieceChar = 'k'
			}
			fmt.Printf("%c ", pieceChar)
		}
		fmt.Println("|")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
	// TODO: Print SideToMove, Castling rights, EnPassant square
}

