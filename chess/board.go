package chess

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

	return b
}